<script>
  import Timer from "./components/Timer.svelte"
  import WorkDay from "./components/WorkDay.svelte";
  import Sarkis from "./assets/images/sarkis-dev.png"
  import CreateProject from "./components/projects/CreateProject.svelte";
  import Projects from "./components/projects/Projects.svelte";
  import Project from "./components/projects/Project.svelte";

  let currentTab = 'Work';

  let projectID = 0

  function updateTab(event) {
    currentTab = event.target.value;
  }

  function handleTabEvent(event){
    projectID = event.detail.project?.projectID
    currentTab = event.detail.tab
  }

  function setTab(tab) {
    currentTab = tab;
  }
</script>

<div class="min-h-screen h-auto flex w-full flex-col justify-between items-center bg-primary p-0 font-body">
  <!-- Dropdown for smaller screens -->
  <div class="sm:hidden">
    <label for="Tab" class="sr-only">Tab</label>
    <select id="Tab" class="w-full rounded-md border-secondary bg-accent" bind:value={currentTab} on:change={updateTab}>
      <option>Work</option>
      <option>Timer</option>
      <option>Projects</option>
    </select>
  </div>

  <!-- Navigation for larger screens -->
  <div class="hidden sm:block bg-secondary w-full p-2">
    <nav class="flex w-full justify-center gap-6 m-auto" aria-label="Tabs">
      <button
        class="shrink-0 rounded-lg p-2 text-sm font-medium text-textSecondary hover:bg-buttonAccentBg hover:text-buttonAccentText"
        on:click={() => setTab('Work')}
        class:bg-buttonPrimaryBg={currentTab === 'Work'}
        class:hover:bg-buttonHoverBg={currentTab === 'Work'}
      >
        Work Day
      </button>

      <button
        class="shrink-0 rounded-lg p-2 text-sm font-medium text-textSecondary hover:bg-buttonAccentBg hover:text-buttonAccentText"
        on:click={() => setTab('Timer')}
        class:bg-buttonPrimaryBg={currentTab === 'Timer'}
        class:hover:bg-buttonHoverBg={currentTab === 'Timer'}
      >
        Timer
      </button>

      <button
        class="shrink-0 rounded-lg p-2 text-sm font-medium text-textSecondary hover:bg-buttonAccentBg hover:text-buttonAccentText"
        on:click={() => setTab('Projects')}
        class:bg-buttonPrimaryBg={currentTab === 'Projects'}
        class:hover:bg-buttonHoverBg={currentTab === 'Projects'}
      >
        Projects
      </button>
    </nav>
  </div>
  <section class="my-2">
    {#if currentTab === 'Work'}
      <WorkDay on:tabEvent={handleTabEvent}></WorkDay>
    {/if}
    {#if currentTab === 'Timer'}
      <Timer></Timer>    
    {/if}
    {#if currentTab === 'Projects'}
      <Projects on:tabEvent={handleTabEvent}></Projects>
    {/if}
    {#if currentTab === 'createProject'}
      <CreateProject on:tabEvent={handleTabEvent}></CreateProject>
    {/if}
    {#if currentTab === 'Project'}
      <Project projectID={projectID}></Project>
    {/if}
  </section>
  
  <footer class="w-full bg-secondary text-textPrimary shadow-sm focus:border-sky-500 focus:ring-sky-500 font-nerd">
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

