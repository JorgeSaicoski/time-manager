<script>
    import { onMount } from 'svelte';
    import { GetDaySummary, UpdateWorkTimeDuration, UpdateBreakTimeDuration } from '../../../wailsjs/go/main/App';
    import Message from '../base/Message.svelte';
    import { format } from 'date-fns';
    import Button from '../base/Button.svelte';
  
    export let data;
  
    let day = null  
    let daySummary = {
      workTimesStarted: [],
      workTimesCrossingDays: [],
      workTimeProjects: [],
      breaks: [],
      brbs: []
    };
  
    let message = "";
    let messageType = "info";
    let currentProject = null;
    let hours = 0;
    let minutes = 0;
    let currentBreakTime = null;
    let currentBreakTimeHours = 0;
    let currentBreakTimeMinutes = 0;
    let isoDay
    let loading = false
  
    onMount(async () => {
        day = data.day;
        loading = true;
        try {
        isoDay = day.toISOString().split('T')[0];
        const summary = await GetDaySummary(isoDay);
        daySummary = summary;
        message = "Day summary loaded successfully!";
        messageType = "info";
        } catch (error) {
        message = `Error loading day summary: ${error.message}`;
        messageType = "error";
        } finally {
        loading = false;
        }
    });
  
    function changeProject(id, duration) {
      currentProject = id;
      const totalSeconds = Math.floor(duration / 1000000000);
      hours = Math.floor(totalSeconds / 3600);
      minutes = Math.floor((totalSeconds % 3600) / 60);
    }
  
    function changeBreakTime(id, duration) {
      currentBreakTime = id;
      const totalSeconds = Math.floor(duration / 1000000000);
      currentBreakTimeHours = Math.floor(totalSeconds / 3600);
      currentBreakTimeMinutes = Math.floor((totalSeconds % 3600) / 60);
    }
  
    async function saveUpdatedDuration(newHours, newMinutes, model) {
        loading = true;
        try {
            const newDurationSeconds = (newHours * 3600 + newMinutes * 60);
            
            if (model === "workTime") {
            await UpdateWorkTimeDuration(currentProject, newDurationSeconds);
            message = "WorkTime duration updated!";
            } else if (model === "breakTime") {
            await UpdateBreakTimeDuration(currentBreakTime, newDurationSeconds);
            message = "BreakTime duration updated!";
            }
            
            isoDay = day.toISOString().split('T')[0];
            const summary = await GetDaySummary(isoDay);
            daySummary = summary;
            
            messageType = "info";
        } catch (error) {
            message = `Error saving updated duration: ${error.message}`;
            messageType = "error";
        } finally {
            loading = false;
        }
    }
  
    function formatTimeToHourAndMinute(startTime) {
      return format(new Date(startTime), 'HH:mm');
    }
  
    function formatDuration(duration) {
      const totalSeconds = Math.floor(duration / 1000000000);
  
      const hoursDuration = Math.floor(totalSeconds / 3600);
      const minutesDuration = Math.floor((totalSeconds % 3600) / 60);
      const secondsDuration = totalSeconds % 60;
  
      return `${hoursDuration}h ${minutesDuration}m ${secondsDuration}s`;
    }
  </script>
  
  <div class="w-full max-w-2xl p-0 space-y-4 bg-secondary rounded-lg shadow-lg text-white">
    {#if loading}
        <p class="bg-accent text-buttonAccentText">Loading...</p>
    {/if}
    <h2>Day Summary</h2>
    {#if message}
      <Message message={message} type={messageType}></Message>
    {/if}
    
    {#if daySummary}
      <div>
        <h2 class="text-xl font-bold">Day Summary for {day}</h2>
        <p>Total Time Worked: {formatDuration(daySummary.totalTime)}</p>
  
        {#if Array.isArray(daySummary.workTimesStarted) && daySummary.workTimesStarted.length > 0}
          <h3 class="font-bold mt-4">Work Times Started</h3>
          <ul>
            {#each daySummary.workTimesStarted as workTime}
              {#if currentProject === workTime.ID}
                <li class="w-full bg-hover">
                  <input type="number" bind:value={hours} min="1" class="text-black w-[35px]" /> hours
                  :
                  <input type="number" bind:value={minutes} min="1" class="text-black w-[35px]" /> minutes
                  <Button label="Save Duration" onClick={() => saveUpdatedDuration(hours, minutes, "workTime")}></Button>
                </li>
              {:else}
                <li>
                  {formatTimeToHourAndMinute(workTime.StartTime)} - {formatDuration(workTime.Duration)}
                  <Button label="Update Duration" onClick={() => changeProject(workTime.ID, workTime.Duration)}></Button>
                </li>
              {/if}
            {/each}
          </ul>
        {/if}
  
        {#if Array.isArray(daySummary.workTimesCrossingDays) && daySummary.workTimesCrossingDays.length > 0}
          <h3 class="font-bold mt-4">Work Times Crossing Days</h3>
          <ul>
            {#each daySummary.workTimesCrossingDays as workTime}
              <li>{formatTimeToHourAndMinute(workTime.StartTime)} - {formatDuration(workTime.Duration)}</li>
            {/each}
          </ul>
        {/if}
  
        {#if Array.isArray(daySummary.breaks) && daySummary.breaks.length > 0}
          <h3 class="font-bold mt-4">Breaks</h3>
          <ul>
            {#each daySummary.breaks as breakTime}
              {#if currentBreakTime === breakTime.ID}
                <li class="w-full bg-hover">
                  <input type="number" bind:value={currentBreakTimeHours} min="1" class="text-black w-[35px]" /> hours
                  :
                  <input type="number" bind:value={currentBreakTimeMinutes} min="1" class="text-black w-[35px]" /> minutes
                  <Button label="Save Duration" onClick={() => saveUpdatedDuration(currentBreakTimeHours, currentBreakTimeMinutes, "breakTime")}></Button>
                </li>
              {:else}
                <li>
                  {formatTimeToHourAndMinute(breakTime.StartTime)} - {formatDuration(breakTime.Duration)}
                  <Button label="Update Duration" onClick={() => changeBreakTime(breakTime.ID, breakTime.Duration)}></Button>
                </li>
              {/if}
            {/each}
          </ul>
        {/if}
  
        {#if Array.isArray(daySummary.workTimeProjects) && daySummary.workTimeProjects.length > 0}
          <h3 class="font-bold mt-4">WorkTime Projects</h3>
          <ul>
            {#each daySummary.workTimeProjects as workTimeProject}
              <li>{workTimeProject.Project.Name}</li>
            {/each}
          </ul>
        {/if}
      </div>
    {/if}
  </div>
  