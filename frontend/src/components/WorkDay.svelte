<script>
  import {
    StartDay,
    TakeBreak,
    StartWorkTime,
    FinishDay,
    EndBreak,
    AssociateProjectToWorkTime,
    GetAllProjects,
    GetUnfinishedWorkTimeProjectWithoutSendingError,
    CreateTask,
    TakeBrb,
    EndBrb,
  } from "../../wailsjs/go/main/App";
  import { onDestroy, onMount, createEventDispatcher } from "svelte";
  import Button from "./base/Button.svelte";
  import Message from "./base/Message.svelte";

  let workDayStarted = false;
  let breakTime = false;
  let brbTime = false;
  let workTime = null;
  let selectedProject = null;
  let projects = [];
  let currentProject = null;
  let message = null;
  let messageType = "info";
  let totalTime = null;
  let timerStart = null;
  let elapsedTime = "00:00:00";
  let interval;
  let intervalName = "Day work";
  let tasks = [];
  let newTaskDescription = "";
  let newTaskDeadline = null;

  const dispatch = createEventDispatcher();

  const checkUnfinishedWorkTimeProject = async () => {
    try {
      const response = await GetUnfinishedWorkTimeProjectWithoutSendingError();
      if (response) {
        selectedProject = response.ProjectID;
        currentProject = associateProject(selectedProject);
      }
    } catch (error) {
      message = +" No unfinished work time project found";
      console.log(error);
    }
  };

  const startWorkDay = async (startWork) => {
    try {
      const response = await StartDay();
      totalTime = response.totalTime;

      if (startWork) {
        await startWorkTime();
      } else {
        message = response.message;
        timerStart = new Date(totalTime.StartTime);
        intervalName = "Day time";
      }
      updateElapsedTime();
      interval = setInterval(updateElapsedTime, 200);
    } catch (err) {
      message = err.message;
    }
    workDayStarted = true;
  };

  const finishWorkDay = async () => {
    try {
      const response = await FinishDay();
      workDayStarted = false;
      message = response;
      messageType = "info";
      totalTime = null;
      workTime = null;
      timerStart = null;
      currentProject = null;
      elapsedTime = "00:00:00";
      interval = null;
      console.log(response);
    } catch (err) {
      message = err.message;
      messageType = "error";
    }
  };

  function updateElapsedTime() {
    if (timerStart) {
      const now = new Date();
      const diff = now - timerStart;

      const hours = String(Math.floor(diff / 3600000)).padStart(2, "0");
      const minutes = String(Math.floor((diff % 3600000) / 60000)).padStart(
        2,
        "0",
      );
      const seconds = String(Math.floor((diff % 60000) / 1000)).padStart(
        2,
        "0",
      );

      elapsedTime = `${hours}:${minutes}:${seconds}`;
    }
  }

  const takeBreak = async () => {
    try {
      let response = await TakeBreak();
      message =
        response +
        " You are on a break. Any interaction with the app will signal the end of your break.";

      breakTime = true;
      timerStart = new Date();
      intervalName = "Break time";
      messageType = "alert";
      currentProject = null;
    } catch (err) {
      console.error("Error while creating break:", err);
      message = "An error occurred while trying to start the break.";
      messageType = "error";
    }
  };
  const endBreak = async () => {
    try {
      const response = await EndBreak();
      breakTime = false;
      workTime = response.workTime;
      messageType = "info";
      message = response.message;
      timerStart = new Date(workTime.StartTime);
      intervalName = "Work time";
      currentProject = null;
    } catch (error) {
      message = err.message;
      messageType = "error";
    }
  };

  function doesWorkTimeExist(workTime) {
    if (workTime.StartTime) {
      const startTime = new Date(workTime.StartTime);
      const currentTime = new Date();
      const durationInSeconds = (currentTime - startTime) / 1000;
      return durationInSeconds > 30;
    }
    return false;
  }

  const startWorkTime = async () => {
    try {
      const response = await StartWorkTime();
      workTime = response.workTime;
      messageType = "info";
      message = doesWorkTimeExist(workTime)
        ? "Welcome back! You're continuing your previous work session."
        : response.message;

      timerStart = new Date(workTime.StartTime);
      intervalName = "Work Session";
    } catch (err) {
      messageType = "error";
      message = `Error: ${err.message}`;
    }
  };

  const fetchProjects = async () => {
    try {
      const response = await GetAllProjects(1, 10, false);
      projects = response.projects;
      checkUnfinishedWorkTimeProject();
    } catch (error) {
      message = "Error loading projects";
      console.log(error);
    }
  };

  const brb = async () => {
    try {
      let response = await TakeBrb();
      message =
        response +
        " You are on a brb. Any interaction with the app will signal the end of your break.";

      brbTime = true;
      timerStart = new Date();
      intervalName = "Brb time";
      messageType = "alert";
      currentProject = null;
    } catch (err) {
      console.error("Error while creating brb:", err);
      message = "An error occurred while trying to start the brb.";
      messageType = "error";
    }
  };

  const brbBack = async () => {
    try {
      const response = await EndBrb();
      brbTime = false;
      workTime = response.workTime;
      messageType = "info";
      message = response.message;
      timerStart = new Date(workTime.StartTime);
      intervalName = "Work time";
    } catch (error) {
      message = err.message;
      messageType = "error";
    }
  };

  function createProject() {
    dispatch("tabEvent", { tab: "createProject" });
  }

  const associateProject = async (projectID) => {
    try {
      const response = await AssociateProjectToWorkTime(projectID);
      currentProject = response.project;
      message = response.message;
      tasks = currentProject.Tasks ? currentProject.Tasks : [];
      if (response.workTimeProject) {
        timerStart = new Date(response.workTimeProject.StartTime);
        intervalName = `Working on ${currentProject.Name} since`;
      }
    } catch (error) {
      message = `Error: ${err.message}`;
    }
  };

  const createTask = async () => {
    if (!newTaskDeadline) {
      message = "Please provide a valid deadline.";
      messageType = "error";
      return;
    }
    try {
      const response = await CreateTask(
        currentProject.ID,
        newTaskDescription,
        newTaskDeadline,
      );
      tasks = [...tasks, response.task];
      message = response.message;
      newTaskDeadline = null;
      newTaskDescription = "";
    } catch (error) {
      message = error.message;
      messageType = "error";
    }
  };

  onDestroy(() => {
    clearInterval(interval);
  });
  onMount(() => {
    startWorkDay(true);
    fetchProjects();
  });
</script>

<div
  class="w-full max-w-2xl p-4 space-y-4 bg-secondary rounded-lg shadow-lg text-white"
>
  <Message {message} type={messageType}></Message>
  {#if !workDayStarted}
    <button
      on:click={() => startWorkDay(false)}
      class="w-full py-3 px-6 bg-green-600 hover:bg-green-700 text-white rounded-md transition-colors duration-200 ease-in-out shadow-md"
    >
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
          <Button label="Finish Break" type="normal" onClick={() => endBreak()}
          ></Button>
        {/if}
        {#if !breakTime}
          {#if brbTime}
            <Button label="BRB Back" type="normal" onClick={() => brbBack()}
            ></Button>
          {:else}
            <Button label="BRB" type="normal" onClick={() => brb()}></Button>
            <Button label="Take Break" type="normal" onClick={() => takeBreak()}
            ></Button>
          {/if}
        {/if}
        <Button
          label="Create Project"
          type="create"
          onClick={() => createProject()}
        ></Button>
      {/if}
      {#if !workTime}
        <Button
          label="Start or return to working"
          type="normal"
          onClick={() => startWorkTime()}
        ></Button>
      {/if}
      <Button label="Finish Day" type="error" onClick={() => finishWorkDay()}
      ></Button>
    </div>
    {#if workTime && !breakTime && !brbTime}
      <div class="w-full">
        <h2 class="text-lg font-bold mb-2">Associate Existing Project</h2>
        <div class="flex gap-4">
          <select
            bind:value={selectedProject}
            class="flex-1 p-3 bg-gray-700 text-black border border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="" disabled selected>Select a project</option>
            {#each projects as project}
              <option value={project.ID}>{project.Name}</option>
            {/each}
          </select>
          <Button
            label="Associate"
            type="normal"
            onClick={() => associateProject(selectedProject)}
            disabled={selectedProject ? false : true}
          ></Button>
        </div>
      </div>
    {/if}

    {#if currentProject}
      <div class="mt-6">
        <h2 class="text-lg font-bold">
          Current Project: {currentProject.Name}
        </h2>
        <ul class="list-disc pl-5 space-y-1 text-gray-300">
          {#each tasks as task, index}
            <li>{task.Description}</li>
          {/each}
        </ul>
        <div class="mt-4">
          <label for="description" class="text-white">Task Description:</label>
          <textarea
            id="description"
            placeholder="Description"
            rows="1"
            type="text"
            bind:value={newTaskDescription}
            class="w-full p-2 bg-primary text-white m-0"
          />
          <div class="w-full bg-primary">
            <label for="date">Deadline:</label>
            <input
              id="date"
              type="date"
              bind:value={newTaskDeadline}
              class="p-2 m-0 bg-gray-800 text-white h-full"
            />
          </div>
          <Button
            label="Add Task"
            onClick={() => {
              createTask();
            }}
            type="normal"
          ></Button>
        </div>
      </div>
    {/if}
  {/if}
</div>
