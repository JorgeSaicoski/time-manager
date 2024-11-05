<script>
  import { onMount } from "svelte";
  import {
    GetUnitsTrackerByDay,
    CreateUnitByDay,
    DeleteResolutionUnit,
  } from "../../../wailsjs/go/main/App";
  import Message from "../base/Message.svelte";
  import Button from "../base/Button.svelte";

  export let data;

  let day = null;
  let units = [];
  let isDayFetched = false;
  let message = null;
  let messageType = "info";
  let identifier = "";

  const fetchDay = async () => {
    try {
      day = data.day;
      const isoDay = day.toISOString().split("T")[0];
      const response = await GetUnitsTrackerByDay(isoDay);
      units = response.units;
      message = response.message;
      isDayFetched = true;
    } catch (error) {
      console.log(error);
      message = error.message;
      messageType = "error";
    }
  };

  const deleteResolutionUnit = async (unitID) => {
    try {
      const response = await DeleteResolutionUnit(unitID);
      if (response.success) {
        message = response.message;
        messageType = "success";
        units = units.filter((unit) => unit.ID !== unitID);
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

  const createResolutionUnit = async () => {
    try {
      const isoDay = day.toISOString().split("T")[0];
      const response = await CreateUnitByDay(identifier, isoDay);
      if (response.success) {
        message = response.message;
        messageType = "success";
        fetchDay();
      } else {
        message = response.message;
        messageType = "error";
      }
    } catch (error) {
      console.log(error);
      message = error.message;
      messageType = "error";
    }
    identifier = "";
  };

  onMount(async () => {
    await fetchDay();
  });
</script>

<div
  class="container mx-auto bg-secondary text-textPrimary p-6 rounded-lg shadow-lg font-nerd w-[1000px]"
>
  <h2>Tracker Summary</h2>
  {#if message}
    <Message {message} type={messageType}></Message>
  {/if}

  {#if isDayFetched}
    <h2>Day fetched {day}</h2>
    <div class="my-5">
      <div class="mt-4">
        <label
          for="identifier"
          class="block text-sm font-medium text-textPrimary"
        >
          New Resolution Unit Identifier:
        </label>
        <input
          id="identifier"
          type="text"
          bind:value={identifier}
          placeholder="Enter a unique identifier (e.g., Tkt #12345)"
          class="w-full p-2 mt-1 bg-primary text-black border border-gray-500 rounded-md"
        />
        <Button
          label="Create Resolution Unit"
          onClick={() => createResolutionUnit()}
        />
      </div>
    </div>
    <div class="my-5">
      <p>Trackers</p>
      <ul>
        {#if Array.isArray(units) && units.length > 0}
          {#each units as unit}
            <li
              class="m-5 border-2 border-primary rounded-md shadow-md bg-buttonPrimaryBg flex flex-row items-center justify-center"
            >
              <p class="w-2/3 text-center">{unit.Identifier}</p>
              <Button
                label="Delete"
                type="error"
                onClick={() => {
                  if (confirm("Are you sure you want to delete this unit?")) {
                    deleteResolutionUnit(unit.ID);
                  }
                }}
              ></Button>
            </li>
          {/each}
        {/if}
      </ul>
    </div>
  {/if}
</div>
