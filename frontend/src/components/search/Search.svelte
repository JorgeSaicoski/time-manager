<script>
  import { Search } from "../../../wailsjs/go/main/App";
  import { format } from "date-fns";
  import Button from "../base/Button.svelte";
  import Message from "../base/Message.svelte";

  let searchTerm = "";
  let results = [];
  let message = null;
  let messageType = "info";
  let loading = false;

  const performSearch = async () => {
    const response = await Search(searchTerm);
    if (!searchTerm.trim()) {
      message = "Please enter a search term.";
      messageType = "error";
      return;
    }

    loading = true;
    message = null;

    try {
      if (response.success) {
        results = response.results;
        if (results.length === 0) {
          message = "No results found.";
          messageType = "info";
        }
      } else {
        message = response.message;
        messageType = "error";
      }
    } catch (error) {
      console.log(error);
      message = error.message;
      messageType = "error";
    } finally {
      console.log(results);
      loading = false;
    }
  };

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return format(date, "yyyy-MM-dd");
  };
</script>

<div
  class="container mx-auto p-6 bg-secondary text-textPrimary rounded-lg shadow-lg"
>
  <h2 class="text-2xl font-bold mb-4">Search</h2>

  {#if message}
    <Message {message} type={messageType}></Message>
  {/if}

  <div class="flex items-center space-x-4 mb-6">
    <input
      type="text"
      bind:value={searchTerm}
      placeholder="Enter search term"
      class="flex-grow p-2 border border-gray-500 rounded-md bg-primary text-black"
    />
    <Button label="Search" onClick={() => performSearch()}></Button>
  </div>

  {#if loading}
    <p>Loading...</p>
  {/if}

  {#if results.length > 0}
    <ul class="space-y-4">
      {#each results as result}
        <li class="p-4 border border-gray-500 rounded-md bg-buttonPrimaryBg">
          {#if result.type === "project"}
            <p>
              <strong>Project:</strong>
              {result.name}
            </p>
          {:else if result.type === "resolution_unit"}
            <p>
              <strong>Resolution Unit:</strong>
              {result.identifier}
            </p>
            <p>
              <strong>Day:</strong>
              {formatDate(result.day)}
            </p>
          {:else if result.type === "task"}
            <p>
              <strong>Task:</strong>
              {result.name}
            </p>
          {/if}
        </li>
      {/each}
    </ul>
  {/if}
</div>
