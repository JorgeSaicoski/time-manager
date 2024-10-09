<script>
  import Timer from "./components/Timer.svelte"

  let currentTab = 'Work';

  function updateTab(event) {
    currentTab = event.target.value;
  }

  function setTab(tab) {
    currentTab = tab;
  }
</script>

<div class="h-full flex w-full flex-col justify-between items-center dark:bg-gray-800 p-2">
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
    Work
  {/if}
  {#if currentTab === 'Timer'}
    <Timer></Timer>    
  {/if}
  {#if currentTab === 'Projects'}
    Projects
  {/if}

</div>
