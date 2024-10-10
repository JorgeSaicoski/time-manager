<script>
    import {StartDay} from "../../wailsjs/go/main/App"

    let workDayStarted = false;
    let newProjectName = '';
    let selectedProject = '';
    let projects = [];
    let tasks = [];
    let currentProject = null;

    function startWorkDay() {
        workDayStarted = true;
    }

    function createProject() {

    }

    function associateProject() {
        
    }

    function createTask(taskName) {
    }
</script>

<div class="w-full max-w-2xl p-4 space-y-4 bg-gray-800 rounded-lg shadow-lg text-white">
    {#if !workDayStarted}
        <button 
        on:click={startWorkDay} 
        class="w-full py-3 px-6 bg-green-600 hover:bg-green-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
        Start Workday
        </button>
    {/if}

    {#if workDayStarted}
        <div class="flex flex-col md:flex-row gap-4">
        <button 
            on:click={takeBreak} 
            class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
            Take Break
        </button>
        <button 
            on:click={brb} 
            class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
            BRB (Not Working in Paid Hour)
        </button>
        <button 
            on:click={createProject} 
            class="flex-1 py-3 px-6 bg-purple-600 hover:bg-purple-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
            Create Project
        </button>
        </div>

        <div class="mt-6">
        <h2 class="text-lg font-bold mb-2">Associate Existing Project</h2>
        <div class="flex gap-4">
            <select 
            bind:value={selectedProject} 
            class="flex-1 p-3 bg-gray-700 text-white border border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">
            <option value="" disabled selected>Select a project</option>
            {#each projects as project}
                <option value={project.name}>{project.name}</option>
            {/each}
            </select>
            <button 
            on:click={associateProject} 
            class="py-3 px-6 bg-teal-500 hover:bg-teal-600 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
            Associate
            </button>
        </div>
        </div>

        {#if currentProject}
        <div class="mt-6">
            <h2 class="text-lg font-bold">Current Project: {currentProject.name}</h2>
            <ul class="list-disc pl-5 space-y-1 text-gray-300">
            {#each currentProject.tasks as task, index}
                <li>{task.name}</li>
            {/each}
            </ul>
            <div class="mt-4">
            <input 
                type="text" 
                placeholder="New task name" 
                class="w-full p-3 bg-gray-700 text-white border border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
            <button 
                on:click={() => createTask(2)} 
                class="mt-3 w-full py-3 px-6 bg-purple-600 hover:bg-purple-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                Add Task
            </button>
            </div>
        </div>
        {/if}
    {/if}
</div>
