<script>
    import {StartDay, TakeBreak, StartWorkTime, FinishDay} from "../../wailsjs/go/main/App"

    let workDayStarted = false;
    let breakTime = false;
    let workTime = null;
    let selectedProject = '';
    let projects = [];
    let currentProject = null;
    let message = ''
    let totalTime = null
    import { onDestroy } from 'svelte';

    let timerStart = null;
    let elapsedTime = "00:00:00";
    let interval;

    const startWorkDay = async ()=>  {
        try {
            const response  = await StartDay();
            message = response.message
            totalTime = response.totalTime
            timerStart = new Date(totalTime.StartTime);
            updateElapsedTime();
            interval = setInterval(updateElapsedTime, 1000);

        } catch(err){
            message = err.message

        }
        workDayStarted = true;
    }

    const finishWorkDay = async ()=>  {
        try {
            const response  = await FinishDay();
            message = response
            totalTime = null
            workTime = null
            timerStart = null;
            elapsedTime = "00:00:00";
            interval = null;

        } catch(err){
            message = err.message

        }
        workDayStarted = true;
    }

    function updateElapsedTime() {
        if (timerStart) {
            const now = new Date();
            const diff = now - timerStart;

            const hours = String(Math.floor(diff / 3600000)).padStart(2, "0");
            const minutes = String(Math.floor((diff % 3600000) / 60000)).padStart(2, "0");
            const seconds = String(Math.floor((diff % 60000) / 1000)).padStart(2, "0");

        elapsedTime = `${hours}:${minutes}:${seconds}`;
        }
    }
 
    const takeBreak = async () => {
        message = await TakeBreak(workTime.ID)
        breakTime = true
        timerStart = new Date(workTime.BreakTime.StartTime);
    }

    const startWorkTime = async () => {
        console.log("Starting work time")
        try{
            const response = await StartWorkTime()
            console.log(response.workTime)
            workTime = response.workTime
            message = response.message
            timerStart = new Date(workTime.StartTime);
        } catch(err){
            message = err.message
        }

    }

    


    function brb() {

    }


    function createProject() {

    }

    function associateProject() {
        
    }

    function createTask(taskName) {
    }

    onDestroy(() => {
        clearInterval(interval);
    });  
</script>

<div class="w-full max-w-2xl p-4 space-y-4 bg-gray-800 rounded-lg shadow-lg text-white">
    {#if !workDayStarted}
        <button 
        on:click={startWorkDay} 
        class="w-full py-3 px-6 bg-green-600 hover:bg-green-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
        Start Workday
        </button>
        <p>{message}</p>
    {/if}

    {#if workDayStarted}
        <div class="flex flex-col md:flex-row gap-4">
            <div class="text-lg font-bold">
                Time Elapsed: {elapsedTime}
            </div>
            <div>{message}</div>
        </div>
        <div class="flex flex-col md:flex-row gap-4">
            {#if workTime}
                <button 
                    on:click={brb} 
                    class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                    BRB (Not Working in Paid Hour)
                </button>
                <button 
                    on:click={takeBreak} 
                    class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                    Take Break
                </button>
                <button 
                    on:click={finishWorkDay} 
                    class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                    Finish Day
                </button>
                <button 
                    on:click={createProject} 
                    class="flex-1 py-3 px-6 bg-purple-600 hover:bg-purple-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                    Create Project
                </button>
                
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
            {/if}
            {#if !workTime}
            <button 
                on:click={startWorkTime} 
                class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                Start or return to working
            </button>
            {/if}
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
