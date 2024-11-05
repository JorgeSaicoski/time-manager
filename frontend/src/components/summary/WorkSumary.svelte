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
    startOfWeek,
    endOfWeek,
    isSameMonth,
    addDays,
  } from "date-fns";
  import Button from "../base/Button.svelte";
  import Message from "../base/Message.svelte";

  let selectedMonth = new Date();
  let daysInMonth = [];
  let weeksInMonth = [];
  let workData = {};
  let message = null;
  let messageType = "info";
  let loading = true;
  let isTrackerMode = false;

  const dispatch = createEventDispatcher();

  const daysOfWeek = [];
  const firstDayOfWeek = startOfWeek(new Date(), { weekStartsOn: 1 });

  for (let i = 0; i < 7; i++) {
    const day = addDays(firstDayOfWeek, i);
    daysOfWeek.push(format(day, "EEE"));
  }

  function goToDay(day) {
    if (isTrackerMode) {
      dispatch("tabEvent", { tab: "TrackSummary", day: { day } });
    } else {
      dispatch("tabEvent", { tab: "DaySummary", day: { day } });
    }
  }

  function updateData() {
    if (isTrackerMode) {
      fetchTrackerData();
    } else {
      fetchMonthSummary();
    }
  }

  const toggleTrackerMode = () => {
    isTrackerMode = !isTrackerMode;
    updateData();
  };

  const getMonthName = (date) => format(date, "MMMM yyyy");

  const changeMonth = (direction) => {
    if (direction === "next") {
      selectedMonth = addMonths(selectedMonth, 1);
    } else if (direction === "prev") {
      selectedMonth = subMonths(selectedMonth, 1);
    }
    updateData();
  };

  function groupDaysIntoWeeks(days) {
    const weeks = [];
    for (let i = 0; i < days.length; i += 7) {
      weeks.push(days.slice(i, i + 7));
    }
    return weeks;
  }

  const fetchTrackerData = async () => {
    if (!isTrackerMode) return;

    const startMonth = startOfMonth(selectedMonth);
    const endMonth = endOfMonth(selectedMonth);

    const start = startOfWeek(startMonth, { weekStartsOn: 1 });
    const end = endOfWeek(endMonth, { weekStartsOn: 1 });

    daysInMonth = eachDayOfInterval({ start, end });
    weeksInMonth = groupDaysIntoWeeks(daysInMonth);

    workData = {};
    loading = true;

    for (let day of daysInMonth) {
      const utcDay = format(day, "yyyy-MM-dd");
      try {
        const response = await GetUnitsTrackerByDay(utcDay);
        workData[utcDay] = response.success ? response.units?.length : 0;
      } catch (err) {
        message = "Error loading tracker data.";
        messageType = "error";
        break;
      }
    }

    loading = false;
  };

  const fetchMonthSummary = async () => {
    const startMonth = startOfMonth(selectedMonth);
    const endMonth = endOfMonth(selectedMonth);

    const start = startOfWeek(startMonth, { weekStartsOn: 1 });
    const end = endOfWeek(endMonth, { weekStartsOn: 1 });

    daysInMonth = eachDayOfInterval({ start, end });
    weeksInMonth = groupDaysIntoWeeks(daysInMonth);

    workData = {};
    loading = true;

    for (let day of daysInMonth) {
      const formattedDate = format(day, "yyyy-MM-dd");
      try {
        const totalTime = await CalculateWorkTimeForDay(
          new Date(`${formattedDate}T00:00:00Z`),
        );
        workData[formattedDate] = totalTime
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
    const formattedDate = format(day, "yyyy-MM-dd");
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
  <Button
    label={isTrackerMode ? "CHANGE FOR HOURS WORKED" : "CHANGE FOR TRACKER"}
    onClick={toggleTrackerMode}
  />

  <div class="flex justify-between items-center mb-6">
    <Button label="PREVIOUS" onClick={() => changeMonth("prev")} />
    <h2 class="text-2xl font-bold">{getMonthName(selectedMonth)}</h2>
    <Button label="NEXT" onClick={() => changeMonth("next")} />
  </div>

  {#if loading}
    <div>Loading data, please wait...</div>
  {:else}
    <!-- Days of the week headers -->
    <div class="grid grid-cols-7 gap-4">
      {#each daysOfWeek as dayOfWeek}
        <div class="text-center font-bold">{dayOfWeek}</div>
      {/each}
    </div>

    <!-- Calendar grid -->
    {#each weeksInMonth as week}
      <div class="grid grid-cols-7 gap-0">
        {#each week as day}
          {#if isSameMonth(day, selectedMonth)}
            <!-- Day belongs to the current month -->
            <button
              on:click={() => goToDay(day)}
              class="text-center p-4 border-2 border-secondary {getWorkTimeForDay(
                day,
              ) > 0
                ? 'bg-hover'
                : 'bg-secondaryAccent'}"
            >
              <p class="text-lg font-bold text-textSecondary">
                {formatDay(day)}
              </p>
              <p
                class={getWorkTimeForDay(day) > 0
                  ? "text-textPrimary"
                  : "text-textDark"}
              >
                {#if isTrackerMode}
                  {getWorkTimeForDay(day)}
                {:else}
                  {getWorkTimeForDay(day)}h
                {/if}
              </p>
            </button>
          {:else}
            <!-- Empty cell or day from another month -->
            <div class="bg-textSecondary border-2 border-secondary"></div>
          {/if}
        {/each}
      </div>
    {/each}
  {/if}
</div>
