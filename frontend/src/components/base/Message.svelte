<script>
    export let message = "";
    export let type = "info"; 
  
    let messageClass = "";
    $: {
      if (type === "info") {
        messageClass = "bg-accent text-primary";
      } else if (type === "error") {
        messageClass = "bg-error text-white";
      }
    }
  </script>
  
  <div class={`p-4 rounded-md ${messageClass}`}>
    <p>{message}</p>
  </div>