<script>
  import { StartTimer, GetStartTimes, RemoveTimer } from "../../wailsjs/go/main/App";
  import { onMount } from "svelte";
  import Message from "./base/Message.svelte";
  import Button from "./base/Button.svelte";

  let timeToTimer = 0;
  let message = "";
  let hours = 1;
  let minutes = 30;
  let alertMessage = "";
  let messageStyle = "info";
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
                      formattedFinishTime: formatFinishTime(timer.FinishTime) // Format the finish time
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

  const removeTimer = async(timer) => {
    try {
      alertMessage = await RemoveTimer(timer);
      messageStyle = "info";
      await fetchActiveTimers();
    } catch (error) {
      console.error("Error:", error);
      alertMessage = "Error removing the timer.";
    }
  };

  const startTimer = async () => {
      timeToTimer = minutes * 60 + hours * 60 * 60;
      if (activeTimers.length >= 3) {
          messageStyle = "error";
          alertMessage = "You cannot set more than 3 timers at the same time.";
          return;
      }

      try {
          const receivedMessage = await StartTimer(timeToTimer, message);
          messageStyle = "info";

          alertMessage = receivedMessage;

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

  const formatFinishTime = (finishTime) => {
      const finish = new Date(finishTime);
      return finish.toLocaleTimeString(); 
  };

  onMount(() => {
      fetchActiveTimers();
  });
</script>

<div class="flex flex-col gap-4 mx-auto p-6 bg-secondary text-textPrimary rounded-lg">
  <div>
    <label for="hours" class="block text-sm font-medium">Hours</label>
    <select
      id="hours"
      class="mt-1 block w-full rounded-md border-textSecondary shadow-sm focus:border-sky-500 focus:ring-textAccent text-textDark"
      bind:value={hours}
    >
      {#each hourOptions as hour}
        <option value={hour}>{hour} {hour === 1 ? 'hour' : 'hours'}</option>
      {/each}
    </select>
  </div>

  <div>
    <label for="minutes" class="block text-sm font-medium">Minutes</label>
    <select
      id="minutes"
      class="mt-1 block w-full rounded-md border-textSecondary shadow-sm focus:border-sky-500 focus:ring-textAccent text-textDark"
      bind:value={minutes}
    >
      {#each minuteOptions as minute}
        <option value={minute}>{minute} {minute === 1 ? 'minute' : 'minutes'}</option>
      {/each}
    </select>
  </div>

  <div>
    <label for="message" class="block text-sm font-medium text-textPrimary">Message</label>
    <textarea
      id="message"
      rows="4"
      class="mt-1 block w-full rounded-md border-textSecondary shadow-sm focus:border-sky-500 focus:ring-textAccent text-textDark"
      placeholder="Enter your message here..."
      bind:value={message}
    ></textarea>
  </div>

  <div class="flex justify-center">
    <Button label="Set Timer" onClick={()=>{startTimer()}}></Button>
  </div>

  {#if alertMessage}
    <Message message={alertMessage} type={messageStyle}></Message>
  {/if}

  <div>
    <h3 class="text-xl font-bold text-gray-200">Active Timers:</h3>
    {#if activeTimers.length === 0}
      <p class="text-gray-400">No active timers.</p>
    {:else}
      <ul class="list-disc list-inside text-gray-200">
        {#each activeTimers as timer}
          <li class="my-2 flex justify-start border-2 rounded-md p-2 bg-secondaryAccent">
            <button
            class="bg-accent rounded-md text-textDark mx-2"
            on:click={()=>{removeTimer(timer.Message)}}
             >Remove</button>
            Timer for "{timer.Message}" - will trigger at: <p class="text-textAccent px-1"> {timer.formattedFinishTime}</p>
          </li>
        {/each}
      </ul>
    {/if}
  </div>
</div>