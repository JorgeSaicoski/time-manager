<script>
  export let label = "Click";
  export let type = 'normal';
  export let onClick;
  export let disable = false

  let buttonClass = "";
    $: {
      if (type === "normal") {
        buttonClass = "bg-blue-600 hover:bg-blue-700 text-white";
      } else if (type === "error") {
        buttonClass = "bg-red-600 hover:bg-red-700 text-white";
      } else{
        buttonClass = "bg-purple-600 hover:bg-purple-700 text-white";
      }
    }
</script>

<button
  on:click={onClick}
  disabled={disable}
  class={`flex-1 py-3 px-6 ${buttonClass} rounded-md shadow-md`}>
  {label}
</button>
