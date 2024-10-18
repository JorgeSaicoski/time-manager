<script>
  export let label = "Click";
  export let type = 'normal';
  export let onClick;
  export let disabled = false

  let buttonClass = "";
    $: {
      if (type === "normal") {
        buttonClass = "bg-buttonPrimaryBg hover:bg-buttonHoverBg text-buttonPrimaryText";
      } else if (type === "error") {
        buttonClass = "bg-buttonErrorBg hover:bg-red-700 text-buttonPrimaryText";
      } else{
        buttonClass = "bg-buttonAccentBg hover:bg-buttonHoverBg text-buttonAccentText hover:bg-buttonPrimaryText";
      }
    }
</script>

<button
  on:click={onClick}
  disabled={disabled}
  class={`flex-1 py-3 px-6 m-2 ${buttonClass} rounded-md shadow-md disabled:bg-slate-500 disabled:text-slate-50 disabled:border-slate-200 disabled:shadow-none`}>
  {label}
</button>
