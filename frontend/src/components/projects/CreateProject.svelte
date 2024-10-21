<script>
    import { CreateProject, AssociateProjectToWorkTime } from "../../../wailsjs/go/main/App";
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher()
    let associateProject = false
    let projectName = "";
    let message = null;
    let project = null;
  
    const handleInput = async () => {
        try {
            const response = await CreateProject(projectName)
            message = response.message
            project = response.project
            projectName = ""
            if (associateProject) {
                await AssociateProjectToWorkTime(project.ID);
                message += ` Project associated with the current work time.`;
            }
            
        } catch (error) {
            message = error.message
            console.error(error);
        }    
    }

    function returnToWorkTime() {
        dispatch('tabEvent', { tab: "Work" });
    }
</script>
  
<div class="w-full max-w-2xl p-4 space-y-4 bg-gray-800 rounded-lg shadow-lg text-white">
    {#if message}
        <div class="w-full max-w-2xl p-4 space-y-0 bg-teal-400 rounded shadow-lg text-black">{message}</div>
    {/if}
    <label for="inputField" class="block text-sm font-medium text-gray-700">
        Enter Project Name
    </label>
    <textarea
      id="inputField"
      rows="1"
      class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-sky-500 focus:ring-sky-500 text-slate-800"
      placeholder="Enter the project name"
      bind:value={projectName}
    ></textarea>
    <div class="flex items-center">
        <input 
          type="checkbox" 
          id="associateProjectCheckbox" 
          class="mr-2"
          bind:checked={associateProject}
        />
        <label for="associateProjectCheckbox" class="text-sm">Associate project with current work time</label>
    </div>
    <div class="flex justify-center">
        <button
          class="px-4 py-2 bg-sky-600 text-white rounded-md hover:bg-sky-700 focus:outline-none focus:ring-2 focus:ring-sky-500"
          on:click={() => handleInput()}
        >
          Create Project
        </button>
      </div>
    {#if project}
        <div class="w-full max-w-2xl p-4 space-y-0 bg-teal-400 rounded shadow-lg text-black">
            Project {project.Name} created with id {project.ID}.
        </div>
        <button 
        on:click={returnToWorkTime} 
        class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
        Return to Work Time
        </button>
    {/if}
</div>