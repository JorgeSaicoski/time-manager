<script>
  export let label = "Click";
  export let type = 'normal';
  export let onClick;
  export let disable = false

  let messageClass = "";
    $: {
      if (type === "normal") {
        messageClass = "secondary";
      } else if (type === "error") {
        messageClass = "error";
      } else{
        messageClass = "accent";
      }
    }
</script>

<button
  on:click={onClick}
  disabled={disable}
  class={`flex-1 py-3 px-6 bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors ease-in-out shadow-md`}>
  {label}
</button>
