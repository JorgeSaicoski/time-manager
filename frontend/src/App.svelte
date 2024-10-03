<script lang="ts">
  import { StartTimer, TestMessage } from '../wailsjs/go/main/App'
  let addTime = 10
  let message = ""
  const startTimer = async () => {
      try {
          await StartTimer(addTime);
      } catch (error) {
          console.error("Error:", error);
          message = "Error"
      }
  };
  const newMessage = async () => {
      try {
          message = await TestMessage(); 
      } catch (error) {
          console.error("Error:", error);
          message = "Error"
      }
  };

</script>

<main>

  <button on:click={startTimer}>
    Start Time
  </button>
  <button on:click={newMessage}>
    New Message
  </button>
  <p>{message}</p>
  
</main>

<style>



</style>
