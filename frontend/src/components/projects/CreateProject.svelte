<script>
    import CreateProject from "../../../wailsjs/go/main/App"

    let projectName = "";
    let message = null;
    let project = null;
  
    function handleInput(event) {
        try {
            const response = CreateProject(project)
            message = response.message
            project = response.project
            
        } catch (error) {
            message = error.message
            console.error(error);
        }    
    }
</script>
  
<div class="p-4 max-w-md mx-auto">
    <label for="inputField" class="block text-sm font-medium text-gray-700">
        Enter Project Name
    </label>
    <input
        id="inputField"
        type="text"
        bind:value={projectName}
        on:input={handleInput}
        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        placeholder="Type something..."
    />
    <p class="mt-2 text-sm text-gray-500">You entered: {inputValue}</p>
</div>