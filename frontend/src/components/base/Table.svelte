<script>
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher()
    export let data = [];
    function goToProject(projectID) {
      dispatch('projectEvent', { projectID: projectID });
    }
  </script>
  
  <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
    <thead class="bg-teal-900 text-gray-400">
      <tr>
        <th class="px-6 py-3 border">Name</th>
        <th class="px-6 py-3 border">ID</th>
        <th class="px-6 py-3 border">Start Time</th>
        <th class="px-6 py-3 border">Closed</th>
      </tr>
    </thead>
    <tbody>
      {#each data as row}
        <tr class="bg-gray-800 odd:bg-gray-700 border-gray-700 hover:bg-gray-50 hover:cursor-pointer" on:click={goToProject(row.ID)}>
          <td class="px-6 py-4 font-medium text-teal-500 whitespace-nowrap dark:text-white border">{row.Name}</td>
          <td class="px-6 py-4 border">{row.ID}</td>
          <td class="px-6 py-4 border">{new Date(row.StartTime).toLocaleDateString()}</td>
          <td class="px-6 py-4 border">{row.Closed ? "Yes" : "No"}</td>
        </tr>
      {/each}
    </tbody>
  </table>