<script>
  import { StartTimer, GetStartTimes } from "../../wailsjs/go/main/App";
  import { onMount } from "svelte";
  import Message from "./base/Message.svelte";

  let timeToTimer = 0;
  let message = "";
  let hours = 1;
  let minutes = 30;
  let alertMessage = "";
  let activeTimers = []; 

  const hourOptions = Array.from({ length: 24 }, (_, i) => i);
  const minuteOptions = Array.from({ length: 60 }, (_, i) => i);

  const fetchActiveTimers = async () => {
        try {
            const timers = await GetStartTimes();
            if (timers) {
                activeTimers = timers.map(timer => {
                    return {
                        ...timer,
                        remainingTime: formatTime(timer.Duration)
                    };
                });
            } else {
                activeTimers = [];
            }
        } catch (error) {
            console.error("Error fetching active timers:", error);
            activeTimers = [];
        }
    };

  const startTimer = async () => {
      timeToTimer = minutes * 60 + hours * 60 * 60;
      if (activeTimers.length >= 3) {
          alertMessage = "You cannot set more than 3 timers at the same time.";
          return;
      }

      try {
          const receivedMessage = await StartTimer(timeToTimer, message);
          alertMessage = `
              ${receivedMessage}
              Timer set for ${hours} hour(s) and ${minutes} minute(s). Message: "${message}"
          `;
          await fetchActiveTimers();
          timeToTimer = 0;
          message = "";
          hours = 0;
          minutes = 0;
      } catch (error) {
          console.error("Error:", error);
          alertMessage = "Error starting the timer.";
      }
  };

  const formatTime = (nanoseconds) => {
        let totalSeconds = nanoseconds / 1e9; 
        let hours = Math.floor(totalSeconds / 3600);
        let minutes = Math.floor((totalSeconds % 3600) / 60);
        return `${hours}h ${minutes}m`;
  };

  onMount(() => {
      fetchActiveTimers();
  });
</script>

<div class="flex flex-col gap-4 max-w-md mx-auto p-6">
  <div>
    <label for="hours" class="block text-sm font-medium text-gray-200">Hours</label>
    <select
      id="hours"
      class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-sky-500 focus:ring-sky-500 text-slate-800"
      bind:value={hours}
    >
      {#each hourOptions as hour}
        <option value={hour}>{hour} {hour === 1 ? 'hour' : 'hours'}</option>
      {/each}
    </select>
  </div>

  <div>
    <label for="minutes" class="block text-sm font-medium text-gray-200">Minutes</label>
    <select
      id="minutes"
      class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-sky-500 focus:ring-sky-500 text-slate-800"
      bind:value={minutes}
    >
      {#each minuteOptions as minute}
        <option value={minute}>{minute} {minute === 1 ? 'minute' : 'minutes'}</option>
      {/each}
    </select>
  </div>

  <div>
    <label for="message" class="block text-sm font-medium text-gray-200">Message</label>
    <textarea
      id="message"
      rows="4"
      class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-sky-500 focus:ring-sky-500 text-slate-800"
      placeholder="Enter your message here..."
      bind:value={message}
    ></textarea>
  </div>

  <div class="flex justify-center">
    <button
      class="px-4 py-2 bg-sky-600 text-white rounded-md hover:bg-sky-700 focus:outline-none focus:ring-2 focus:ring-sky-500"
      on:click={startTimer}
    >
      Set Timer
    </button>
  </div>

  <Message message={alertMessage} type="info"></Message>

  <div>
    <h3 class="text-xl font-bold text-gray-200">Active Timers:</h3>
    {#if activeTimers.length === 0}
      <p class="text-gray-400">No active timers.</p>
    {:else}
      <ul class="list-disc list-inside text-gray-200">
        {#each activeTimers as timer}
          <li>
            Timer for "{timer.Message}" - {timer.remainingTime} remaining.
          </li>
        {/each}
      </ul>
    {/if}
  </div>
</div>
