<script>
  import { onMount } from "svelte";
  import { GetUnitsTrackerByDay } from "../../../wailsjs/go/main/App";
  import Message from "../base/Message.svelte";

  export let data;

  let day = null;
  let units = [];
  let isDayFetched = false;
  let message = null;
  let messageType = "info";

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

  onMount(async () => {
    console.log("mount");
    await fetchDay();
  });
</script>

<div>
  <h2>Tracker Summary</h2>
  {#if message}
    <Message {message} type={messageType}></Message>
  {/if}
  {#if isDayFetched}
    <h2>Day fetched {day}</h2>
    <p>Trackers</p>
    <ul>
      {#each units as unit}
        <li>
          {unit.Identifier}
        </li>
      {/each}
    </ul>
  {/if}
</div>
