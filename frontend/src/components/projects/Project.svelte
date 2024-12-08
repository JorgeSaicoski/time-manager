<script>
  import { onMount } from "svelte";
  import {
    GetProjectByID,
    CreateTask,
    ChangeProjectClose,
    CalculateAndSaveProjectCost,
    UpdateProjectName,
    DeleteTask,
  } from "../../../wailsjs/go/main/App";
  import Message from "../base/Message.svelte";
  import Button from "../base/Button.svelte";

  export let projectID = null;

  let message = null;
  let messageType = "info";
  let project = null;
  let newTaskDescription = "";
  let newTaskDeadline = null;
  let hourCost = 0;
  let totalMinutes = 0;
  let hours = 0;
  let minutes = 0;

  const changeCost = async () => {
    try {
      hourCost = parseInt(hourCost, 10);
      if (isNaN(hourCost)) {
        throw new Error("Invalid cost. Please provide a number.");
      }
      const response = await CalculateAndSaveProjectCost(projectID, hourCost);
      message = response.message;
      hourCost = response.cost.HourCost;
      messageType = "info";
      project.Cost = response.cost;
    } catch (err) {
      console.log(err);
      message = err.message;
      messageType = "error";
    }
  };

  const updateProjectName = async () => {
    try {
      const response = await UpdateProjectName(projectID, project.Name);
      message = response.message;
      messageType = "info";
      project = response.project;
    } catch (err) {
      message = err.message;
      messageType = "error";
    }
  };

  const removeTask = async (taskID) => {
    try {
      const response = await DeleteTask(taskID);
      if (response) {
        message = response.message;
        messageType = "success";
        project.Tasks = project.Tasks.filter((task) => task.ID !== taskID);
      } else {
        message = response.message;
        messageType = "error";
      }
    } catch (error) {
      console.log(error);
      message = error.message;
      messageType = "error";
    }
  };

  const addTask = async () => {
    if (!newTaskDeadline) {
      message = "Please provide a valid deadline.";
      messageType = "error";
      return;
    }
    try {
      const response = await CreateTask(
        project.ID,
        newTaskDescription,
        newTaskDeadline,
      );
      project.Tasks = [...project.Tasks, response.task];
      message = response.message;
      newTaskDeadline = null;
      newTaskDescription = "";
    } catch (error) {
      message = error.message;
      messageType = "error";
    }
  };

  const changeCloseProject = async () => {
    try {
      const response = await ChangeProjectClose(projectID);
      message = response.message;
      messageType = "info";
      project = response.project;
    } catch (err) {
      message = err.message;
      messageType = "error";
    }
  };

  const findProject = async () => {
    try {
      const response = await GetProjectByID(projectID);
      message = response.message;
      messageType = "info";
      project = response.project;
      if (project?.Duration) {
        totalMinutes = project.Duration / 60000000000;
        hours = Math.floor(totalMinutes / 60);
        minutes = Math.floor(totalMinutes % 60);
      } else {
        totalMinutes = 0;
        hours = 0;
        minutes = 0;
      }
    } catch (err) {
      message = err.message;
      messageType = "error";
    }
  };
  onMount(async () => {
    await findProject();
    hourCost = project?.Cost?.HourCost ? project.Cost.HourCost : 10;
    project.Tasks = project.Tasks ? project.Tasks : [];
    changeCost();
  });
</script>

<div
  class="container mx-auto bg-secondary text-textPrimary p-6 rounded-lg shadow-lg font-nerd"
>
  {#if message}
    <Message {message} type={messageType}></Message>
  {/if}
  {#if project}
    <!-- Project Details -->
    <h1 class="text-3xl font-extrabold mb-6 text-center text-textSecondary">
      {project.Name}
    </h1>

    <div class="grid grid-cols-2 gap-6 mb-6">
      <div>
        <h2 class="text-xl font-bold text-textSecondary mb-2">
          Project Details
        </h2>
        <ul>
          <li class="mb-2">ID: {project.ID}</li>
          <li class="mb-2">
            Start Time: {new Date(project.StartTime).toLocaleString()}
          </li>
          <li class="mb-2">
            Duration:
            {#if hours > 0}
              {hours} {hours === 1 ? "hour" : "hours"}
            {/if}
            {#if minutes > 0}
              {minutes} {minutes === 1 ? "minute" : "minutes"}
            {/if}
            {#if hours === 0 && minutes === 0}
              Less than a minute
            {/if}
          </li>
          <li class="mb-2">Closed: {project.Closed ? "Yes" : "No"}</li>
          <li class="mb-2">
            Cost:
            {#if project.Cost?.HourCost > 0}
              ${(
                project.Cost.HourCost * hours +
                (project.Cost.HourCost / 60) * minutes
              ).toFixed(2)}
            {:else}
              "Not Set"
            {/if}
          </li>
        </ul>
      </div>

      <!-- Edit Project Name -->
      <div>
        <h2 class="text-xl font-bold text-textSecondary mb-2">
          Edit Project Name
        </h2>
        <input
          type="text"
          bind:value={project.Name}
          placeholder="Enter new name"
          class="w-full p-2 rounded-lg bg-textPrimary text-textDark"
        />
        <Button
          label="Save Name"
          type="normal"
          onClick={() => updateProjectName()}
        ></Button>
        <!-- Close Project -->
        <Button
          label={project.Closed ? "Re Open Project" : "Close Project"}
          type="error"
          onClick={() => changeCloseProject()}
        ></Button>
      </div>
    </div>

    <!-- Tasks Section -->
    <div class="mb-6">
      <h2 class="text-xl font-bold text-textSecondary mb-4">Tasks</h2>

      <!-- Task List -->
      <ul class="mb-4">
        {#each project.Tasks as task}
          <li class="bg-secondaryAccent p-4 rounded-lg mb-2">
            <div class="flex justify-between items-center">
              <div class="w-2/3">
                <p>Task ID: {task.ID}</p>
                <p>Description: {task.Description}</p>
                <p>Deadline: {new Date(task.Deadline).toLocaleString()}</p>
                <p>Closed: {task.Closed ? "Yes" : "No"}</p>
              </div>
              <Button
                label="Remove Task"
                type="error"
                onClick={() => {
                  if (confirm("Are you sure you want to delete this unit?")) {
                    removeTask(task.ID);
                  }
                }}
              ></Button>
            </div>
          </li>
        {/each}
      </ul>

      <!-- Add Task -->
      <div class="bg-secondaryAccent rounded-2xl text-textDark">
        <label for="description text-primary">Task Description:</label>
        <textarea
          id="description"
          placeholder="Description"
          rows="5"
          type="text"
          bind:value={newTaskDescription}
          class="w-full p-2 bg-textPrimary m-0"
        />
        <div class="w-full bg-textPrimary">
          <label for="date">Deadline:</label>
          <input
            id="date"
            type="date"
            bind:value={newTaskDeadline}
            class="p-2 m-0 h-full"
          />
        </div>
        <Button label="Add Task" type="normal" onClick={() => addTask()}
        ></Button>
      </div>
    </div>

    <!-- Add Cost Section -->
    <div>
      <h2 class="text-xl font-bold text-textSecondary mb-2">
        Set Project Cost Per Hour
      </h2>
      <div
        class="w-full rounded-lg p-0 bg-secondary text-textPrimary flex justify-center"
      >
        <span>$</span>
        <input
          class="bg-secondary h-full text-textPrimary m-0"
          bind:value={hourCost}
          type="number"
          min="0"
        />
        <span>per hour</span>
      </div>
      <Button
        label="Save Cost"
        onClick={() => {
          changeCost();
        }}
      ></Button>
    </div>
  {/if}
</div>
