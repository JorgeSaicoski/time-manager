<script>
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
  let category = "tkt";
  let newUnitIdentifier = "";

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
      This tracker allows you to monitor daily tasks or cases, like customer
      support tickets. You can start a tracker for today, categorize your tasks,
      and add specific cases or tickets as resolution units.
    </p>
    <Button
      label="Start or Recover Tracker for Today"
      on:click={startOrRecoverTracker}
    />
  {/if}

  <!-- Display the tracker options once started -->
  {#if tracker}
    <h2 class="text-lg font-bold mt-4">Today's Tracker</h2>

    <Message {message} type={messageType} />

    <!-- Update Tracker Category -->
    <div class="mt-4">
      <label for="category" class="block text-sm font-medium text-textPrimary"
        >Category:</label
      >
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
      <Button label="Update Category" on:click={updateCategory} />
    </div>

    <!-- Create Resolution Unit -->
    <div class="mt-4">
      <label for="identifier" class="block text-sm font-medium text-textPrimary"
        >New Resolution Unit Identifier:</label
      >
      <input
        id="identifier"
        type="text"
        bind:value={newUnitIdentifier}
        placeholder="Enter a unique identifier"
        class="w-full p-2 mt-1 bg-primary text-black border border-gray-500 rounded-md"
      />
      <Button label="Create Resolution Unit" on:click={createResolutionUnit} />
    </div>
  {/if}
</div>

<style>
  .tracker-container {
    background-color: var(--secondary-color);
    color: var(--textPrimary);
  }
</style>
