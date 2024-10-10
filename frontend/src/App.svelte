<script>
  import Timer from "./components/Timer.svelte"
  import WorkDay from "./components/WorkDay.svelte";
  import Sarkis from "./assets/images/sarkis-dev.png"

  let currentTab = 'Work';

  function updateTab(event) {
    currentTab = event.target.value;
  }

  function setTab(tab) {
    currentTab = tab;
  }
</script>

<div class="h-full flex w-full flex-col justify-between items-center dark:bg-gray-800 p-2 font-body">
  <!-- Dropdown for smaller screens -->
  <div class="sm:hidden">
    <label for="Tab" class="sr-only">Tab</label>
    <select id="Tab" class="w-full rounded-md border-gray-200" bind:value={currentTab} on:change={updateTab}>
      <option>Work</option>
      <option>Timer</option>
      <option>Projects</option>
    </select>
  </div>

  <!-- Navigation for larger screens -->
  <div class="hidden sm:block">
    <nav class="flex gap-6" aria-label="Tabs">
      <button
        class="shrink-0 rounded-lg p-2 text-sm font-medium text-gray-500 hover:bg-gray-50 hover:text-gray-700"
        on:click={() => setTab('Work')}
        class:text-sky-600={currentTab === 'Work'}
        class:bg-sky-100={currentTab === 'Work'}
      >
        Work Day
      </button>

      <button
        class="shrink-0 rounded-lg p-2 text-sm font-medium text-gray-500 hover:bg-gray-50 hover:text-gray-700"
        on:click={() => setTab('Timer')}
        class:text-sky-600={currentTab === 'Timer'}
        class:bg-sky-100={currentTab === 'Timer'}
      >
        Timer
      </button>

      <button
        class="shrink-0 rounded-lg p-2 text-sm font-medium text-gray-500 hover:bg-gray-50 hover:text-gray-700"
        on:click={() => setTab('Projects')}
        class:text-sky-600={currentTab === 'Projects'}
        class:bg-sky-100={currentTab === 'Projects'}
      >
        Projects
      </button>
    </nav>
  </div>
  {#if currentTab === 'Work'}
    <WorkDay></WorkDay>
  {/if}
  {#if currentTab === 'Timer'}
    <Timer></Timer>    
  {/if}
  {#if currentTab === 'Projects'}
    Projects
  {/if}
  <footer class="w-full bg-black shadow-sm focus:border-sky-500 focus:ring-sky-500 text-teal-800 font-nerd">
    <div class="mx-auto flex flex-col items-center space-y-4">
      <p>
        Developed by: 
      </p>
      <img src={Sarkis} alt="Sarkis DEV logo" class="w-32 h-auto" />
      <p >
        Contact: <a href="mailto:jorge@sarkis.dev" >jorge@sarkis.dev</a>
      </p>
    </div>
  </footer>
</div>

