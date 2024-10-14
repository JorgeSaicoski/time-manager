<script>
    export let label = "Click";
    export let disabled = false;
    export let onClick;
  </script>
  
  <button
    on:click={onClick}
    class="bg-primary text-white font-semibold px-6 py-3 rounded-lg shadow-md hover:bg-hover disabled:opacity-50"
    disabled={disabled}
  >
    {label}
  </button>