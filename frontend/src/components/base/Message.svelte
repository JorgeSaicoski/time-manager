<script>
    export let message = "";
    export let type = "info"; 
  
    let messageClass = "";
    $: {
      if (type === "info") {
        messageClass = "bg-secondary text-white";
      } else if (type === "error") {
        messageClass = "bg-error text-white";
      } else{
        messageClass = "bg-accent text-secondary"
      }
    }
  </script>
  
  <div class={`p-4 rounded-md ${messageClass}`}>
    <p>{message}</p>
  </div>