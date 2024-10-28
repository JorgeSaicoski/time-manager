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

  // Variables for edit functionality
  let currentEdit = { type: null, id: null, hours: 0, minutes: 0 };

  onMount(async () => {
    await loadCurrentActiveTimers();
  });

  async function loadCurrentActiveTimers() {
    loading = true;
    try {
      const timers = await GetCurrentActiveTimers();
      currentTimers = timers;
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
</script>

<div class="w-full max-w-2xl p-4 bg-secondary rounded-lg shadow-lg text-white">
  {#if loading}
    <p class="bg-accent text-buttonAccentText">Loading...</p>
  {/if}

  <h2>Currently Active Timers</h2>
  {#if message}
    <Message {message} type={messageType}></Message>
  {/if}

  <ul>
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
          <p>Duration: {currentTimers.workTime.Duration}</p>
          <Button
            label="Edit"
            onClick={() => startEdit("workTime", currentTimers.workTime)}
          ></Button>
        {/if}
      </li>
    {/if}

    {#if currentTimers.breakTime}
      <li>
        <h3>Break Time</h3>
        {#if currentEdit.type === "breakTime" && currentEdit.id === currentTimers.breakTime.ID}
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
          <p>Duration: {currentTimers.breakTime.Duration}</p>
          <Button
            label="Edit"
            onClick={() => startEdit("breakTime", currentTimers.breakTime)}
          ></Button>
        {/if}
      </li>
    {/if}

    {#if currentTimers.brb}
      <li>
        <h3>BRB Time</h3>
        {#if currentEdit.type === "brbTime" && currentEdit.id === currentTimers.brb.ID}
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
          <p>Duration: {currentTimers.brb.Duration}</p>
          <Button
            label="Edit"
            onClick={() => startEdit("brbTime", currentTimers.brb)}
          ></Button>
        {/if}
      </li>
    {/if}
  </ul>
</div>
