<script>
    import {StartTimer} from "../../wailsjs/go/main/App"
    
    let timeToTimer = 0
    let message = ""
    let hours = 1;
    let minutes = 30;
    let alertMessage = ""

    const hourOptions = Array.from({ length: 24 }, (_, i) => i);
    const minuteOptions = Array.from({ length: 60 }, (_, i) => i);

    const startTimer = async () => {
        timeToTimer = (minutes*60) + (hours*60*60)
        try {
             const receivedMessage = await StartTimer(timeToTimer, message);
             alertMessage = `
                ${receivedMessage}
                Timer set for ${hours} hour(s) and ${minutes} minute(s). Message: "${message}"
             `
        } catch (error) {
            console.error("Error:", error);
            message = "Error starting the timer.";
        }
    };

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
      on:click={() => startTimer()}
    >
      Set Timer
    </button>
  </div>
  <div>
      {alertMessage}
  </div>
</div>
  
