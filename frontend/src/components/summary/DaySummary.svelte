<script>
  import { onMount } from 'svelte';
  import { GetDaySummary } from '../../../wailsjs/go/main/App';
  import Message from '../base/Message.svelte';
  import {  format } from 'date-fns';

  export let data;  

  let day = null  
  let daySummary = {
    workTimesStarted:[],
    workTimesCrossingDays:[],
    projects:[],
    breaks:[],
    brbs:[]

  };
  let message = "";
  let messageType = "info";
  
  onMount(async () => {
    day = data.day
    try {
        const isoDay = day.toISOString().split('T')[0]; 
        const summary = await GetDaySummary(isoDay);
        console.log(summary)
        daySummary = summary;
        message = "Day summary loaded successfully!";
        messageType = "info";
    } catch (error) {
        message = `Error loading day summary: ${error.message}`;
        messageType = "error";
    }
  });



  function formatDuration(duration) {
      const hours = Math.floor(duration / 3600);
      const minutes = Math.floor((duration % 3600) / 60);
      const seconds = Math.floor(duration % 60);
      return `${hours}h ${minutes}m ${seconds}s`;
  }

</script>


<div class="w-full max-w-2xl p-4 space-y-4 bg-secondary rounded-lg shadow-lg text-white">
  <h2>Day Summary</h2>
  {#if message}
    <Message message={message} type={messageType}></Message>
  {/if}
  
  {#if daySummary}
      <div>
          <h2 class="text-xl font-bold">Day Summary for {day}</h2>
          <p>Total Time Worked: {formatDuration(daySummary.totalTime)}</p>

          <h3 class="font-bold mt-4">Work Times Started</h3>
          <ul>
              {#each daySummary.workTimesStarted as workTime}
                  <li>{workTime.StartTime} - {formatDuration(workTime.Duration)}</li>
              {/each}
          </ul>

          <h3 class="font-bold mt-4">Work Times Crossing Days</h3>
          <ul>
              {#each daySummary.workTimesCrossingDays as workTime}
                  <li>{workTime.StartTime} - {formatDuration(workTime.Duration)}</li>
              {/each}
          </ul>

          <h3 class="font-bold mt-4">Projects</h3>
          <ul>
              {#each daySummary.projects as project}
                  <li>{project.Name}</li>
              {/each}
          </ul>

          <h3 class="font-bold mt-4">Breaks</h3>
          <ul>
              {#each daySummary.breaks as breakTime}
                  <li>{breakTime.StartTime} - {formatDuration(breakTime.Duration)}</li>
              {/each}
          </ul>

          <h3 class="font-bold mt-4">BRBs</h3>
          <ul>
              {#each daySummary.brbs as brb}
                  <li>{brb.StartTime} - {formatDuration(brb.Duration)}</li>
              {/each}
          </ul>
      </div>
  {/if}
</div>
