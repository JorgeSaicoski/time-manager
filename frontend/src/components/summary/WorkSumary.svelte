<script>
    import { onMount } from 'svelte';
    import { GetDaySummary } from '../../../wailsjs/go/main/App'; 
    import Message from '../base/Message.svelte';
    import { eachDayOfInterval, format, startOfMonth, endOfMonth } from 'date-fns'; 

    let selectedMonth = new Date();  
    let daysInMonth = [];
    let workData = {};
    let message = null;
    let messageType = "info";

    const fetchMonthSummary = async () => {
        const start = startOfMonth(selectedMonth);
        const end = endOfMonth(selectedMonth);
        daysInMonth = eachDayOfInterval({ start, end });

        workData = {};  // Reset work data

        for (let day of daysInMonth) {
            try {
                const response = await GetDaySummary(day);
                Ge
                workData[format(day, 'yyyy-MM-dd')] = response.totalTime ? (response.totalTime / (60 * 60 * 1e9)).toFixed(1) : 0;  // Convert nanoseconds to hours
            } catch (err) {
                message = "Error loading work data.";
                messageType = "error";
                break;
            }
        }
    }

    onMount(() => {
        fetchMonthSummary();
    });

    const formatDay = (day) => format(day, 'd');
    const getWorkTimeForDay = (day) => workData[format(day, 'yyyy-MM-dd')] || 0;
</script>

<div class="container mx-auto bg-secondary text-textPrimary p-6 rounded-lg shadow-lg font-nerd">
    {#if message}
        <Message message={message} type={messageType}></Message>
    {/if}

    <div class="grid grid-cols-7 gap-4">
        {#each daysInMonth as day}
            <div class="text-center p-4 rounded-lg {getWorkTimeForDay(day) > 0 ? 'bg-blue-100' : 'bg-secondaryAccent'}">
                <p class="text-lg font-bold">{formatDay(day)}</p>
                <p class={getWorkTimeForDay(day) > 0 ? 'text-blue-600' : 'text-gray-400'}>
                    {getWorkTimeForDay(day)}h
                </p>
            </div>
        {/each}
    </div>
</div>