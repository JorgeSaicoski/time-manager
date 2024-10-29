<script>
  import { onMount } from "svelte";
  import Button from "./base/Button.svelte";
  import {
    GetOrCreateTodayResolutionTracker,
    UpdateResolutionTrackerCategory,
    CreateResolutionUnit,
  } from "../../wailsjs/go/main/App";
  import Message from "./base/Message.svelte";

  let tracker = null;
  let message = "";
  let messageType = "info";
  let showExplanation = true;
  let category = "tkt"; // Default category
  let newUnitIdentifier = ""; // For creating a new resolution unit

  // Fetches or initializes the tracker for today
  async function startOrRecoverTracker() {
    const response = await GetOrCreateTodayResolutionTracker();
    if (response.success) {
      tracker = response.tracker;
      showExplanation = false;
      message = "Today's tracker started or recovered successfully!";
      messageType = "info";
    } else {
      message = "Failed to start or recover today's tracker.";
      messageType = "error";
    }
  }

  // Updates the category of the current tracker
  async function updateCategory() {
    if (tracker) {
      const response = await UpdateResolutionTrackerCategory(
        tracker.ID,
        category,
      );
      if (response.success) {
        message = "Tracker category updated successfully!";
        messageType = "info";
      } else {
        message = "Failed to update tracker category.";
        messageType = "error";
      }
    }
  }

  // Creates a new resolution unit in the current tracker
  async function createResolutionUnit() {
    if (tracker && newUnitIdentifier) {
      const response = await CreateResolutionUnit(newUnitIdentifier);
      if (response.success) {
        message = "Resolution unit created successfully!";
        messageType = "info";
        newUnitIdentifier = ""; // Reset input field
      } else {
        message = "Failed to create resolution unit.";
        messageType = "error";
      }
    } else {
      message = "Please enter an identifier for the new resolution unit.";
      messageType = "warning";
    }
  }
</script>

<div
  class="tracker-container max-w-2xl p-4 bg-secondary rounded-lg shadow-lg text-white"
>
  <!-- Display an explanation if the tracker hasn't started -->
  {#if showExplanation}
    <h2 class="text-lg font-bold">What is this?</h2>
    <p class="mt-2">
      This tracker allows you to monitor and record daily cases or tickets
      resolved. It serves as a tool to help you track your daily performance,
      identify trends, or provide evidence of completed tasks when needed.
      Whether you're working with support tickets, customer cases, or other task
      types, this tracker can help you stay on top of your daily workload.
    </p>
    <p class="mt-2">
      For example, if you need to prove the number of cases resolved each day or
      check your progress on tasks, this tool will record each case or task as a
      "resolution unit," giving you a record of your daily achievements.
    </p>
    <Button
      label="Start or Recover Tracker for Today"
      onClick={startOrRecoverTracker}
    />
  {/if}

  <!-- Display the tracker options once started -->
  {#if tracker}
    <h2 class="text-lg font-bold mt-4">Today's Tracker</h2>

    <Message {message} type={messageType} />

    <!-- Update Tracker Category -->
    <div class="mt-4">
      <label for="category" class="block text-sm font-medium text-textPrimary">
        Category:
      </label>
      <select
        id="category"
        bind:value={category}
        class="w-full p-2 bg-primary text-black border border-gray-500 rounded-md"
      >
        <option value="tkt">Ticket</option>
        <option value="support">Support</option>
        <option value="development">Development</option>
        <!-- Add more categories as needed -->
      </select>
      <Button label="Update Category" onClick={updateCategory} />
    </div>

    <!-- Create Resolution Unit -->
    <div class="mt-4">
      <label
        for="identifier"
        class="block text-sm font-medium text-textPrimary"
      >
        New Resolution Unit Identifier:
      </label>
      <p class="text-sm text-gray-300 mb-2">
        Enter a unique identifier to easily recognize this task or case.
        Examples: "Tkt #12345," "Client A support call," or "Issue with X
        software." Use details that help you remember the specific case or
        ticket.
      </p>
      <input
        id="identifier"
        type="text"
        bind:value={newUnitIdentifier}
        placeholder="Enter a unique identifier (e.g., Tkt #12345)"
        class="w-full p-2 mt-1 bg-primary text-black border border-gray-500 rounded-md"
      />
      <Button label="Create Resolution Unit" onClick={createResolutionUnit} />
    </div>
  {/if}
</div>
