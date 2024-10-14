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
    let projectID = 0

    const dispatch = createEventDispatcher()

    function goToProject(event) {
      projectID = event.detail.projectID
      dispatch('tabEvent', { tab: "Project", project:{projectID} });
    }


  
    const fetchProjects = async (page) => {
      try {
        const response = await GetAllProjects(page, pageSize);
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
  </script>
  
  <div class="container mx-auto bg-gray-900 text-white p-6 rounded-lg shadow-lg font-nerd">
    <h1 class="text-3xl font-extrabold mb-6 text-center text-teal-300">Projects</h1>
  
    {#if message}
      <Message message={message} type="error" />
    {/if}
  
    <Table data={projects} on:projectEvent={goToProject}/>
  
    <div class="flex justify-between mt-6">
      <Button label="Previous" onClick={prevPage} disable={currentPage === 1} />
      <Button label="Next" onClick={nextPage} disable={currentPage === totalPages} />
    </div>
  </div>
