<script>
  import Timer from "./components/Timer.svelte";
  import WorkDay from "./components/WorkDay.svelte";
  import Sarkis from "./assets/images/sarkis-dev.png";
  import Donation from "./assets/images/donation.png";
  import CreateProject from "./components/projects/CreateProject.svelte";
  import Projects from "./components/projects/Projects.svelte";
  import Project from "./components/projects/Project.svelte";
  import WorkSumary from "./components/summary/WorkSumary.svelte";
  import DaySummary from "./components/summary/DaySummary.svelte";
  import Currently from "./components/Currently.svelte";
  import Tracker from "./components/Tracker.svelte";
  import TrackSummary from "./components/summary/TrackSummary.svelte";
  import Search from "./components/search/Search.svelte";

  let currentTab = "Currently";

  let projectID = 0;
  let day = null;
  const navItem = [
    "Currently",
    "Work",
    "Timer",
    "Projects",
    "Summary",
    "Tracker",
    "Search",
  ];

  function updateTab(event) {
    currentTab = event.target.value;
  }

  function handleTabEvent(event) {
    projectID = event.detail.project?.projectID;
    if (event.detail.day) {
      day = event.detail.day;
    }
    currentTab = event.detail.tab;
  }

  function setTab(tab) {
    currentTab = tab;
  }
</script>

<div
  class="min-h-screen h-auto flex w-full flex-col justify-between items-center bg-primary p-0 font-body"
>
  <!-- Dropdown for smaller screens -->
  <div class="sm:hidden">
    <label for="Tab" class="sr-only">Tab</label>
    <select
      id="Tab"
      class="w-full rounded-md border-secondary bg-accent"
      bind:value={currentTab}
      on:change={updateTab}
    >
      <option>Work</option>
      <option>Timer</option>
      <option>Projects</option>
    </select>
  </div>

  <!-- Navigation for larger screens -->
  <div class="hidden sm:block bg-secondary w-full p-2">
    <nav class="flex w-full justify-center gap-6 m-auto" aria-label="Tabs">
      {#each navItem as tab}
        <button
          class="shrink-0 rounded-lg p-2 text-sm font-medium text-textSecondary hover:bg-buttonAccentBg hover:text-buttonAccentText"
          on:click={() => setTab(tab)}
          class:bg-buttonPrimaryBg={currentTab === tab}
          class:hover:bg-buttonHoverBg={currentTab === tab}
        >
          {tab}
        </button>
      {/each}
    </nav>
  </div>
  <section class="my-2">
    {#if currentTab === "Currently"}
      <Currently></Currently>
    {/if}
    {#if currentTab === "Work"}
      <WorkDay on:tabEvent={handleTabEvent}></WorkDay>
    {/if}
    {#if currentTab === "Timer"}
      <Timer></Timer>
    {/if}
    {#if currentTab === "Projects"}
      <Projects on:tabEvent={handleTabEvent}></Projects>
    {/if}
    {#if currentTab === "createProject"}
      <CreateProject on:tabEvent={handleTabEvent}></CreateProject>
    {/if}
    {#if currentTab === "Project"}
      <Project {projectID}></Project>
    {/if}
    {#if currentTab === "Summary"}
      <WorkSumary on:tabEvent={handleTabEvent}></WorkSumary>
    {/if}
    {#if currentTab === "DaySummary"}
      <DaySummary data={day}></DaySummary>
    {/if}
    {#if currentTab === "Tracker"}
      <Tracker></Tracker>
    {/if}
    {#if currentTab === "TrackSummary"}
      <TrackSummary data={day}></TrackSummary>
    {/if}
    {#if currentTab === "Search"}
      <Search on:tabEvent={handleTabEvent}></Search>
    {/if}
  </section>

  <footer
    class="w-full bg-secondary text-textPrimary shadow-sm font-nerd flex flex-row justify-between p-6 items-center"
  >
    <!-- Contact Information -->
    <div class="flex flex-col items-center space-y-4">
      <h3>Developed by:</h3>
      <img src={Sarkis} alt="Sarkis DEV logo" class="w-32 h-auto" />
      <p>
        Contact: <a href="mailto:jorge@sarkis.dev">jorge@sarkis.dev</a>
      </p>
      <p>
        Code: <a href="https://github.com/JorgeSaicoski/time-manager"
          >github.com/JorgeSaicoski/time-manager</a
        >
      </p>
    </div>

    <!-- Donation Section -->
    <div class="flex flex-col items-center space-y-4">
      <h3>Support Our Work</h3>
      <p>Link for donations:</p>
      <img src={Donation} alt="Donation QR Code" class="w-32 h-auto" />
      <p
        class="bg-gray-700 text-white p-2 rounded-lg font-mono text-sm text-center"
      >
        Bitcoin Address: bc1qew0sxw8ph248mqs8t5s3hceu02es4pjpx8737w
      </p>
    </div>
  </footer>
</div>
