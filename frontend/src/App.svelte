<script>
  import { StartTimer, StartWorkTime, TakeBreak, EndBreak } from '../wailsjs/go/main/App';
  import { EventsOn } from '../wailsjs/runtime/runtime';

  let addTime = 10;
  let message = "";
  let currentWorkTime = null;
  let activeWorkTime = false;

  // Start Work Time
  const startWork = async () => {
    try {
      const response = await StartWorkTime();
      message = response[0];
      currentWorkTime = response[1];
      activeWorkTime = true;
    } catch (error) {
      console.error("Error:", error);
      message = "Error starting work time.";
    }
  };

  // Take a Break
  const takeBreak = async () => {
    try {
      if (currentWorkTime) {
        message = await TakeBreak(currentWorkTime.ID);
        activeWorkTime = false;
      } else {
        message = "Please start a work session first.";
      }
    } catch (error) {
      console.error("Error:", error);
      message = "Error taking a break.";
    }
  };

  // End a Break
  const endBreak = async () => {
    try {
      message = await EndBreak();
      activeWorkTime = true;
    } catch (error) {
      console.error("Error:", error);
      message = "Error ending the break.";
    }
  };

  // Start Timer
  const startTimer = async () => {
    try {
      await StartTimer(addTime, "Focus time!");
    } catch (error) {
      console.error("Error:", error);
      message = "Error starting the timer.";
    }
  };

  // Listen to Timer Events
  EventsOn("timerFinished", (data) => {
    message = data;
  });
</script>

<div class="flex flex-col items-center p-4">
  <h1 class="text-2xl font-bold mb-4">Workday Manager</h1>


  <!-- Conditional UI based on activeWorkTime -->
  {#if !activeWorkTime}
    <!-- Show Start Work Session button if no work time is active -->
    <div class="mb-6">
      <button
        on:click={startWork}
        class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600">
        Start Work Session
      </button>
      
      <button
        on:click={endBreak}
        class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 ml-2">
        End Break
      </button>
    </div>
  {:else}
    <!-- Show Break Controls if work time is active -->
    <div class="mb-6">
      <button
        on:click={takeBreak}
        class="bg-yellow-500 text-white px-4 py-2 rounded hover:bg-yellow-600">
        Start Break
      </button>
      <p class="mt-2">{message}</p>
    </div>

    <!-- Add Project button (to be implemented with backend) -->
    <div class="mb-6">
      <button
        class="bg-purple-500 text-white px-4 py-2 rounded hover:bg-purple-600">
        Work on a Project
      </button>
    </div>
  {/if}

  <!-- Set a Timer -->
  <div class="mb-6">
    <button
      on:click={startTimer}
      class="bg-purple-500 text-white px-4 py-2 rounded hover:bg-purple-600">
      Set 10-second Timer
    </button>
  </div>
  <p class="mt-2">{message}</p>
</div>