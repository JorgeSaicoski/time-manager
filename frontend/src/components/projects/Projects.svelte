<script>
  import { onMount, createEventDispatcher } from 'svelte';
  import { GetAllProjects } from '../../../wailsjs/go/main/App';
  import Table from '../base/Table.svelte';
  import Button from '../base/Button.svelte';
  import Message from '../base/Message.svelte';
  
  let projects = [];
  let currentPage = 1;
  let pageSize = 5;
  let totalPages = 0;
  let message = null;
  let projectID = 0;

  let closedFilter = null; 
  let orderBy = 'updated_at'; 
  let orderDirection = 'desc'; 

  const dispatch = createEventDispatcher();

  function goToProject(event) {
    projectID = event.detail.projectID;
    dispatch('tabEvent', { tab: "Project", project: { projectID } });
  }

  const fetchProjects = async (page) => {
    try {
      const response = await GetAllProjects(page, pageSize, closedFilter, orderBy, orderDirection);
      projects = response.projects;
      currentPage = response.currentPage;
      totalPages = Math.ceil(response.total / response.itemsPerPage);
    } catch (error) {
      message = "Error loading projects";
    }
  };


  const nextPage = () => {
     if (currentPage < totalPages) {
      currentPage++;
      fetchProjects(currentPage);
    }
  };

  const prevPage = () => {
    if (currentPage > 1) {
      currentPage--;
      fetchProjects(currentPage);
    }
  };

  onMount(() => {
    fetchProjects(currentPage);
  });

  const setClosedFilter = (value) => {
    closedFilter = value;
    fetchProjects(currentPage);
  };

  const setOrderBy = (value) => {
    orderBy = value;
    fetchProjects(currentPage);
  };

  const setOrderDirection = (value) => {
    orderDirection = value;
    fetchProjects(currentPage);
  };
</script>

  
<div class="container mx-auto bg-secondary text-textPrimary p-6 rounded-lg shadow-lg font-nerd">
<h1 class="text-3xl font-extrabold mb-6 text-center text-textSecondary">Projects</h1>

{#if message}
  <Message message={message} type="error" />
{/if}

<!-- Filters and sorting options -->
<div class="flex justify-between mb-6">
  <div>
    <!-- Closed filter -->
    <label for="close">Show Closed:</label>
    <div class="w-full max-w-sm min-w-[200px]">      
      <div class="relative">
        <select
        class="w-full bg-transparent placeholder:text-textSecondary text-text-Primary text-sm border border-slate-200 rounded pl-3 pr-8 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md appearance-none cursor-pointer"
        on:change="{(e) => setClosedFilter(e.target.value === 'null' ? null : e.target.value === 'true')}"
        >
          <option value="null">All</option>
          <option value="false">Open</option>
          <option value="true">Closed</option>
        </select>
      </div>
    </div>
  </div>

  <div>
    <!-- Order by filter -->
    <label for="order">Order By:</label>
    <div class="w-full max-w-sm min-w-[200px]"> 
      <select 
      class="w-full bg-transparent placeholder:text-textSecondary text-text-Primary text-sm border border-slate-200 rounded pl-3 pr-8 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md appearance-none cursor-pointer"
      id="order" 
      on:change="{(e) => setOrderBy(e.target.value)}"
      >
        <option value="updated_at">Updated At</option>
        <option value="duration">Duration</option>
      </select>
    </div>
  </div>

  <div>
    <!-- Order direction filter -->
    <label for="direction">Order Direction:</label>
    <div class="w-full max-w-sm min-w-[200px]"> 
      <select 
      id="direction" on:change="{(e) => setOrderDirection(e.target.value)}"
      class="w-full bg-transparent placeholder:text-textSecondary text-text-Primary text-sm border border-slate-200 rounded pl-3 pr-8 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md appearance-none cursor-pointer"
      >
        <option value="desc">Descending</option>
        <option value="asc">Ascending</option>
      </select>
    </div>
  </div>
</div>

<!-- Projects Table -->
<Table data={projects} on:projectEvent={goToProject} />

<!-- Pagination -->
<div class="flex justify-between mt-6">
  <Button label="Previous" onClick={prevPage} disabled={currentPage === 1} />
  <Button label="Next" onClick={nextPage} disabled={currentPage === totalPages} />
</div>
</div>

