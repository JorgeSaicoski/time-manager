<script>
  import { StartTimer, StartTotalTime, StartWorkTime, TakeBreak, EndBreak } from '../wailsjs/go/main/App'
  import { EventsOn } from '../wailsjs/runtime/runtime'
  
  let addTime = 10;
  let message = "";
  let currentWorkTime = {
    projects: [],
    brb: {}
  }
  
  const startTimer = async () => {
    try {
      await StartTimer(addTime, "Focus time!");
    } catch (error) {
      console.error("Error:", error);
      message = "Error starting the timer.";
    }
  };

  const startTotalTime = async () => {
    try {
      message = await StartTotalTime();
    } catch (error) {
      console.error("Error:", error);
      message = "Error starting the workday.";
    }
  };

  const startWork = async () => {
    try {
      message = await StartWorkTime();
    } catch (error) {
      console.error("Error:", error);
      message = "Error starting work time.";
    }
  };

  const takeBreak = async () => {
    try {
      message = await TakeBreak(1); 
    } catch (error) {
      console.error("Error:", error);
      message = "Error taking a break.";
    }
  };

  const endBreak = async () => {
    try {
      message = await EndBreak();
    } catch (error) {
      console.error("Error:", error);
      message = "Error ending the break.";
    }
  };

  EventsOn("timerFinished", (data) => {
    message = data;
  });
</script>

<div class="flex flex-col items-center p-4">
  <h1 class="text-2xl font-bold mb-4">Workday Manager</h1>

  <!-- Start Workday -->
  <div class="mb-6">
    <button
      on:click={startTotalTime}
      class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600">
      Start Workday
    </button>
    <p class="mt-2">{message}</p>
  </div>

  <!-- Start Work Time -->
  <div class="mb-6">
    <button
      on:click={startWork}
      class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600">
      Start Work Session
    </button>
    <p class="mt-2">{message}</p>
  </div>

  <!-- Break Controls -->
  <div class="mb-6">
    <button
      on:click={takeBreak}
      class="bg-yellow-500 text-white px-4 py-2 rounded hover:bg-yellow-600">
      Start Break
    </button>
    <button
      on:click={endBreak}
      class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 ml-2">
      End Break
    </button>
    <p class="mt-2">{message}</p>
  </div>

  <!-- Set a Timer -->
  <div class="mb-6">
    <button
      on:click={startTimer}
      class="bg-purple-500 text-white px-4 py-2 rounded hover:bg-purple-600">
      Set 10-second Timer
    </button>
    <p class="mt-2">{message}</p>
  </div>
</div>
