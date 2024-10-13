<script>
    import { onMount } from 'svelte';
    import { GetAllProjects } from '../../../wailsjs/go/main/App';
  import App from '../../App.svelte';
  
    let projects = [];
    let currentPage = 1;
    let pageSize = 5;
    let totalPages = 0;
  
    const fetchProjects = async (page) => {
        const response = await GetAllProjects(page,pageSize)
        projects = response.projects      
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
  
  <div class="container mx-auto bg-gray-900 text-white p-6 rounded-lg shadow-lg font-nerd">
    <h1 class="text-3xl font-extrabold mb-6 text-center text-teal-300">Projects</h1>
  
    <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
      <thead class=" bg-gray-700 text-gray-400"> 
        <tr>
        <th class="px-6 py-3 border">Name</th>
          <th class="px-6 py-3 border">ID</th>
          <th class="px-6 py-3 border">Start Time</th>
          <th class="px-6 py-3 border">Closed</th>
        </tr>
      </thead>
      <tbody>
        {#each projects as project}
          <tr class="bg-gray-800 border-gray-700 hover:bg-gray-50"> 
            <td class="px-6 py-4 font-medium text-teal-500 whitespace-nowrap dark:text-white border">{project.Name}</td>
            <td class="px-6 py-4 border">{project.ID}</td>
            <td class="px-6 py-4 border">{new Date(project.StartTime).toLocaleDateString()}</td>
            <td class="px-6 py-4 border">{project.Closed ? "Yes" : "No"}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  
    <!-- Pagination Controls -->
    <div class="flex justify-between mt-6">
      <button
        on:click={prevPage}
        class="bg-teal-600 text-white font-semibold px-6 py-3 rounded-lg shadow-md hover:shadow-lg disabled:opacity-50"
        disabled={currentPage === 1}
      >
        Previous
      </button>
  
      <button
        on:click={nextPage}
        class="bg-teal-600 text-white font-semibold px-6 py-3 rounded-lg shadow-md hover:shadow-lg disabled:opacity-50"
        disabled={currentPage === totalPages}
      >
        Next
      </button>
    </div>
</div>
