<script lang="ts">
  import { StartTimer, StartTotalTime } from '../wailsjs/go/main/App'
  import { EventsOn } from '../wailsjs/runtime/runtime'
  let addTime = 10
  let message = ""
  const startTimer = async () => {
      try {
          await StartTimer(addTime, "New Message");
      } catch (error) {
          console.error("Error:", error);
          message = "Error"
      }
  };
  const startTotalTime = async () => {
      try {
          message = await StartTotalTime(); 
      } catch (error) {
          console.error("Error:", error);
          message = "Error"
      }
  };
  EventsOn("timerFinished", (data: string) => {
      message = data;
  });


</script>

<main>

  <button on:click={startTimer}>
    Start Time
  </button>
  <button on:click={startTotalTime}>
    New Message
  </button>
  <p>{message}</p>
  
</main>

<style>



</style>
