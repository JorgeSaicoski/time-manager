<script>
  import { onMount, createEventDispatcher } from "svelte";
  import {
    CalculateWorkTimeForDay,
    GetUnitsTrackerByDay,
  } from "../../../wailsjs/go/main/App";
  import {
    eachDayOfInterval,
    format,
    startOfMonth,
    endOfMonth,
    addMonths,
    subMonths,
  } from "date-fns";
  import Button from "../base/Button.svelte";
  import Message from "../base/Message.svelte";

  let selectedMonth = new Date();
  let daysInMonth = [];
  let workData = {};
  let message = null;
  let messageType = "info";
  let loading = true;
  let isTrackerMode = false;

  const dispatch = createEventDispatcher();

  function goToDay(day) {
    dispatch("tabEvent", { tab: "DaySummary", day: { day } });
  }

  const toggleTrackerMode = () => {
    isTrackerMode = !isTrackerMode;
    fetchTrackerData();
  };

  const getMonthName = (date) => format(date, "MMMM yyyy");

  const changeMonth = (direction) => {
    if (direction === "next") {
      selectedMonth = addMonths(selectedMonth, 1);
    } else if (direction === "prev") {
      selectedMonth = subMonths(selectedMonth, 1);
    }
    fetchMonthSummary();
  };

  const fetchTrackerData = async () => {
    if (!isTrackerMode) return;

    const start = startOfMonth(selectedMonth);
    const end = endOfMonth(selectedMonth);
    daysInMonth = eachDayOfInterval({ start, end });

    workData = {};
    loading = true;

    for (let day of daysInMonth) {
      const utcDay = format(day, "yyyy-MM-dd");
      try {
        const response = await GetUnitsTrackerByDay(utcDay);
        console.log(`Response for ${utcDay}: `, response);
        workData[utcDay] = response.success ? response.units?.length : 0;
      } catch (err) {
        console.log(`Error loading tracker data for ${utcDay}`, err);
        message = "Error loading tracker data.";
        messageType = "error";
        break;
      }
    }

    loading = false;
  };

  const fetchMonthSummary = async () => {
    const start = startOfMonth(selectedMonth);
    const end = endOfMonth(selectedMonth);
    daysInMonth = eachDayOfInterval({ start, end });

    workData = {};
    loading = true;

    for (let day of daysInMonth) {
      const utcDay = new Date(
        Date.UTC(day.getFullYear(), day.getMonth(), day.getDate()),
      );
      try {
        const totalTime = await CalculateWorkTimeForDay(utcDay);
        workData[format(utcDay, "yyyy-MM-dd")] = totalTime
          ? (totalTime / (1e9 * 3600)).toFixed(1)
          : 0;
      } catch (err) {
        message = "Error loading work data.";
        messageType = "error";
        break;
      }
    }

    loading = false;
  };

  onMount(() => {
    fetchMonthSummary();
  });

  const formatDay = (day) => format(day, "d");

  const getWorkTimeForDay = (day) => {
    const formattedDate = format(day, "yyyy-MM-dd"); // Use pre-formatted date as key

    return workData[formattedDate] || 0;
  };
</script>

<!-- Calendar component -->
<div
  class="container mx-auto bg-secondary text-textPrimary p-6 rounded-lg shadow-lg font-nerd w-[1000px]"
>
  {#if message}
    <Message {message} type={messageType}></Message>
  {/if}
  <Button label="Toggle Tracker Mode" onClick={() => toggleTrackerMode()} />

  <div class="flex justify-between items-center mb-6">
    <Button label="Previous" onClick={() => changeMonth("prev")} />
    <h2 class="text-2xl font-bold">{getMonthName(selectedMonth)}</h2>
    <Button label="Next" onClick={() => changeMonth("next")} />
  </div>

  {#if loading}
    <div>Loading data, please wait...</div>
  {:else}
    <div class="grid grid-cols-7 gap-4">
      {#each daysInMonth as day}
        <button
          on:click={goToDay(day)}
          class="text-center p-4 rounded-lg {getWorkTimeForDay(day) > 0
            ? 'bg-hover'
            : 'bg-secondaryAccent'}"
        >
          <p class="text-lg font-bold text-textSecondary">{formatDay(day)}</p>
          <p
            class={getWorkTimeForDay(day) > 0
              ? "text-textPrimary"
              : "text-textDark"}
          >
            {#if isTrackerMode}
              <p>{getWorkTimeForDay(day)}</p>
            {:else}
              <p>{getWorkTimeForDay(day)}h</p>
            {/if}
          </p>
        </button>
      {/each}
    </div>
  {/if}
</div>
