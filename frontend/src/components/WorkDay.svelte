<script>
    import {StartDay, TakeBreak, StartWorkTime, FinishDay, EndBreak} from "../../wailsjs/go/main/App"
    import { onDestroy, onMount, createEventDispatcher } from 'svelte';
  import Button from "./base/Button.svelte";

    let workDayStarted = false;
    let breakTime = false;
    let workTime = null;
    let selectedProject = '';
    let projects = [];
    let currentProject = null;
    let message = ''
    let totalTime = null
    let timerStart = null;
    let elapsedTime = "00:00:00";
    let interval;
    let intervalName= "Day work"

    const dispatch = createEventDispatcher()

    const startWorkDay = async (startWork)=>  {
        console.log("Start Work Day")
        try {
            const response  = await StartDay();
            totalTime = response.totalTime

            if (startWork){
                await startWorkTime()
            }else{
                message = response.message
                timerStart = new Date(totalTime.StartTime)
            }
            updateElapsedTime();
            interval = setInterval(updateElapsedTime, 300);
            intervalName = "Day time"

        } catch(err){
            message = err.message

        }
        workDayStarted = true;
    }

    const finishWorkDay = async ()=>  {
        try {
            const response  = await FinishDay();
            workDayStarted = false
            message = response
            totalTime = null
            workTime = null
            timerStart = null;
            elapsedTime = "00:00:00";
            interval = null;
            console.log(response)

        } catch(err){
            message = err.message

        }
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
        timerStart = new Date();
        intervalName = "Break time"
    }
    const endBreak = async () => {
        try {
            const response = await EndBreak()
            console.log(response.workTime)
            breakTime = false
            workTime = response.workTime
            message = response.message
            timerStart = new Date(workTime.StartTime);
            intervalName = "Work time"
        } catch (error) {
            message = err.message
        }
    }

    function doesWorkTimeExist(workTime){
        console.log(workTime.StartTime)
        if (workTime.StartTime){
            const startTime = new Date(workTime.StartTime);
            const currentTime = new Date();
            const durationInSeconds = (currentTime - startTime) / 1000;
            return durationInSeconds > 30
        }
        return false
    }

    const startWorkTime = async () => {
        try {
            const response = await StartWorkTime();
            workTime = response.workTime;
            
            message = doesWorkTimeExist(workTime) 
                ? "Welcome back! You're continuing your previous work session."
                : response.message;
            
            timerStart = new Date(workTime.StartTime);
            intervalName = "Work Session";
        } catch (err) {
            message = `Error: ${err.message}`;
        }
    };

    


    function brb() {

    }


    function createProject() {
        dispatch('tabEvent', { tab: "createProject" });
    }

    function associateProject() {
        
    }

    function createTask(taskName) {
    }

    onDestroy(() => {
        clearInterval(interval);
    });  
    onMount(()=>{
        console.log("mount")
        startWorkDay(true)
    })
</script>


<div class="w-full max-w-2xl p-4 space-y-4 bg-gray-800 rounded-lg shadow-lg text-white">
    <div class="w-full max-w-2xl p-4 space-y-0 bg-teal-400 rounded shadow-lg text-black">{message}</div>
    {#if !workDayStarted}
        <button 
        on:click={() => startWorkDay(false)}
        class="w-full py-3 px-6 bg-green-600 hover:bg-green-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
        Start Workday
        </button>
    {/if}

    {#if workDayStarted}
        <div class="flex flex-col md:flex-row gap-4">
            <div class="text-lg font-bold">
                {intervalName}: {elapsedTime}
            </div>

        </div>
        <div class="flex flex-col md:flex-row gap-4">
            {#if workTime}
                {#if breakTime}
                    <button 
                        on:click={endBreak} 
                        class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                        Back to work
                    </button>
                {/if}
                {#if !breakTime}
                    <button 
                        on:click={brb} 
                        class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                        BRB (Not Working in Paid Hour)
                    </button>
                    <Button label="Take Break" type="normal" onClick={() => takeBreak()} ></Button>
                    <button 
                        
                        class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                        Take Break
                    </button>
                {/if}
                
                <button 
                    on:click={createProject} 
                    class="flex-1 py-3 px-6 bg-purple-600 hover:bg-purple-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                    Create Project
                </button>            
            {/if}
            {#if !workTime}
            <button 
                on:click={startWorkTime} 
                class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                Start or return to working
            </button>
            {/if}
            <button 
                on:click={finishWorkDay} 
                class="flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md">
                Finish Day
            </button>
        </div>
        {#if workTime}
            <div class="w-full">
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
            <Button message="Add Task" type="normal"></Button>
            </div>
        </div>
        {/if}
    {/if}
</div>
