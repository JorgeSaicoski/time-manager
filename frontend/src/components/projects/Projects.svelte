<script>
    import { onMount } from 'svelte';
    import { GetAllProjects } from '../../../wailsjs/go/main/App';
  
    let projects = [];
    let currentPage = 1;
    let pageSize = 5;
    let totalPages = 0;
  
    const fetchProjects = async (page) => {
        const response = await GetAllProjects(currentPage,pageSize)
        console.log(response)
        projects = response.projects;
        currentPage = response.currentPage;
        totalPages = Math.ceil(response.total / response.itemsPerPage);
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
  
    // Initial load
    onMount(() => {
      fetchProjects(currentPage);
    });
  </script>
  
  <div class="container mx-auto">
    <h1 class="text-2xl font-bold mb-4">Projects</h1>
  
    <table class="min-w-full bg-teal-600 border">
      <thead>
        <tr>
          <th class="px-4 py-2 border">ID</th>
          <th class="px-4 py-2 border">Name</th>
          <th class="px-4 py-2 border">Start Time</th>
          <th class="px-4 py-2 border">Closed</th>
        </tr>
      </thead>
      <tbody>
        {#each projects as project}
          <tr>
            <td class="px-4 py-2 border">{project.id}</td>
            <td class="px-4 py-2 border">{project.name}</td>
            <td class="px-4 py-2 border">{new Date(project.startTime).toLocaleDateString()}</td>
            <td class="px-4 py-2 border">{project.closed ? "Yes" : "No"}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  
    <!-- Pagination Controls -->
    <div class="flex justify-between mt-4">
      <button
        on:click={prevPage}
        class="bg-blue-500 text-white px-4 py-2 rounded disabled:opacity-50"
        disabled={currentPage === 1}
      >
        Previous
      </button>
  
      <button
        on:click={nextPage}
        class="bg-blue-500 text-white px-4 py-2 rounded disabled:opacity-50"
        disabled={currentPage === totalPages}
      >
        Next
      </button>
    </div>
  </div>
  