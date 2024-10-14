<script>
  export let label = "Click";
  export let type = 'normal';
  export let onClick;
  export let disabled = false

  let buttonClass = "";
    $: {
      if (type === "normal") {
        buttonClass = "bg-blue-600 hover:bg-blue-800 text-white";
      } else if (type === "error") {
        buttonClass = "bg-red-600 hover:bg-red-700 text-white";
      } else{
        buttonClass = "bg-purple-600 hover:bg-purple-700 text-white";
      }
    }
</script>

<button
  on:click={onClick}
  disabled={disabled}
  class={`flex-1 py-3 px-6 mx-5 ${buttonClass} rounded-md shadow-md disabled:bg-slate-500 disabled:text-slate-50 disabled:border-slate-200 disabled:shadow-none`}>
  {label}
</button>
