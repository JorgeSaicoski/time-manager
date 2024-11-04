<script>
  import { onMount } from "svelte";
  import {
    GetCurrentActiveTimers,
    UpdateWorkTimeDuration,
    UpdateBreakTimeDuration,
    UpdateBrbDuration,
  } from "../../wailsjs/go/main/App";
  import Message from "./base/Message.svelte";
  import Button from "./base/Button.svelte";

  let currentTimers = {
    workTime: null,
    breakTime: null,
    brb: null,
    totalTime: null,
  };
  let message = "";
  let messageType = "info";
  let loading = false;
  let showGuide = true;

  let currentEdit = { type: null, id: null, hours: 0, minutes: 0 };

  onMount(async () => {
    await loadCurrentActiveTimers();
  });

  async function loadCurrentActiveTimers() {
    loading = true;
    try {
      const timers = await GetCurrentActiveTimers();
      currentTimers = timers;
      showGuide = !(timers.workTime || timers.breakTime || timers.brb);
      message = "Active timers loaded successfully!";
      messageType = "info";
    } catch (error) {
      message = `Error loading active timers: ${error.message}`;
      messageType = "error";
    } finally {
      loading = false;
    }
  }

  function startEdit(type, timer) {
    currentEdit.type = type;
    currentEdit.id = timer.ID;
    const totalSeconds = Math.floor(timer.Duration / 1000000000);
    currentEdit.hours = Math.floor(totalSeconds / 3600);
    currentEdit.minutes = Math.floor((totalSeconds % 3600) / 60);
  }

  async function saveEdit() {
    loading = true;
    try {
      const newDurationSeconds =
        currentEdit.hours * 3600 + currentEdit.minutes * 60;

      if (currentEdit.type === "workTime") {
        await UpdateWorkTimeDuration(currentEdit.id, newDurationSeconds);
        message = "WorkTime duration updated!";
      } else if (currentEdit.type === "breakTime") {
        await UpdateBreakTimeDuration(currentEdit.id, newDurationSeconds);
        message = "BreakTime duration updated!";
      } else if (currentEdit.type === "brbTime") {
        await UpdateBrbDuration(currentEdit.id, newDurationSeconds);
        message = "BRB duration updated!";
      }

      await loadCurrentActiveTimers(); // Refresh timers after update
      messageType = "info";
    } catch (error) {
      message = `Error saving updated duration: ${error.message}`;
      messageType = "error";
    } finally {
      loading = false;
      currentEdit.type = null;
      currentEdit.id = null;
    }
  }

  function formatDuration(duration, createdAt) {
    let totalSeconds = duration ? Math.floor(duration / 1000000000) : 0;

    if (totalSeconds === 0 && createdAt) {
      const start = new Date(createdAt);
      const now = new Date();
      totalSeconds = Math.floor((now - start) / 1000);
    }

    const hours = Math.floor(totalSeconds / 3600);
    const minutes = Math.floor((totalSeconds % 3600) / 60);
    const seconds = totalSeconds % 60;
    return `${hours}h ${minutes}m ${seconds}s`;
  }
</script>

<div class="w-full max-w-2xl p-4 bg-secondary rounded-lg shadow-lg text-white">
  {#if loading}
    <p class="bg-accent text-buttonAccentText">Loading...</p>
  {/if}

  <h2>Currently Active Timers</h2>
  {#if message}
    <Message {message} type={messageType}></Message>
  {/if}

  <!-- Guide Section if No Timers Are Active -->
  {#if showGuide}
    <div class="bg-secondaryAccent p-4 rounded-md mt-4">
      <h3 class="text-lg font-bold mb-2">Timer Categories Explained</h3>
      <p>
        <strong>Break:</strong> This is non-remunerated time when you're not working
        at all. Use this timer for any true pauses away from work.
      </p>

      <p>
        <strong>BRB:</strong> Use BRB if you’re on a paid pause, temporarily out
        (like for a bathroom break), or waiting on tasks. Time logged here is not
        applied to a specific project but can be distributed across projects later
        as needed.
      </p>

      <p>
        <strong>TotalTime:</strong> Typically represents the entire workday. However,
        you can have multiple TotalTime sessions if you're dividing your day among
        different work contexts, such as switching between multiple companies, freelance
        work, or taking a break and resuming later in the day.
      </p>

      <p>
        <strong>WorkTime:</strong> A continuous work session that can include multiple
        projects. This timer runs as long as you're actively working. It only pauses
        if you start a Break or BRB or end the day. Projects can be switched within
        the same WorkTime session without stopping it.
      </p>
      <p>
        <strong>Projects:</strong> Manage tasks and costs associated with different
        projects.
      </p>
      <p>
        <strong>Summary:</strong> View a calendar with daily work hours. Click on
        a specific day for details or corrections.
      </p>
      <p>
        <strong>Timer:</strong> A simple timer to set reminders, allowing up to three
        timers simultaneously.
      </p>
    </div>
  {/if}

  <!-- Currently Active Timers Section -->
  <ul>
    {#if currentTimers.totalTime}
      <li>
        <h3>Total Time</h3>
        <p>
          Duration: {formatDuration(
            currentTimers.totalTime.Duration,
            currentTimers.totalTime.CreatedAt,
          )}
        </p>
      </li>
    {/if}

    {#if currentTimers.workTime}
      <li>
        <h3>Work Time</h3>
        {#if currentEdit.type === "workTime" && currentEdit.id === currentTimers.workTime.ID}
          <input
            type="number"
            bind:value={currentEdit.hours}
            min="0"
            class="text-black w-12"
          />
          hours :
          <input
            type="number"
            bind:value={currentEdit.minutes}
            min="0"
            class="text-black w-12"
          />
          minutes
          <Button label="Save" onClick={saveEdit}></Button>
        {:else}
          <p>
            Duration: {formatDuration(
              currentTimers.workTime.Duration,
              currentTimers.workTime.CreatedAt,
            )}
          </p>
          <Button
            label="Edit"
            onClick={() => startEdit("workTime", currentTimers.workTime)}
          ></Button>
        {/if}
      </li>
    {/if}
  </ul>
</div>
