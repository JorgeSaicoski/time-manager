<script>
    import { onMount } from 'svelte';
    import { CalculateWorkTimeForDay } from '../../../wailsjs/go/main/App';
    import { eachDayOfInterval, format, startOfMonth, endOfMonth, addMonths, subMonths } from 'date-fns';
    import Button from '../base/Button.svelte'; 
    import Message from '../base/Message.svelte';

    let selectedMonth = new Date(); 
    let daysInMonth = [];
    let workData = {};
    let message = null;
    let messageType = "info";


    const getMonthName = (date) => format(date, 'MMMM yyyy');

    const changeMonth = (direction) => {
        if (direction === 'next') {
            selectedMonth = addMonths(selectedMonth, 1);
        } else if (direction === 'prev') {
            selectedMonth = subMonths(selectedMonth, 1);
        }
        fetchMonthSummary();
    }

    const fetchMonthSummary = async () => {
        const start = startOfMonth(selectedMonth);
        const end = endOfMonth(selectedMonth);
        daysInMonth = eachDayOfInterval({ start, end });

        workData = {};

        for (let day of daysInMonth) {
            console.log(day)
            try {
                console.log("try from try catch")
                const totalTime = await CalculateWorkTimeForDay(day);  
                workData[format(day, 'yyyy-MM-dd')] = totalTime ? (totalTime / (60 * 60 * 1e9)).toFixed(1) : 0; 
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

    <div class="flex justify-between items-center mb-6">
        <Button label="Previous" onClick={() => changeMonth('prev')} />
        <h2 class="text-2xl font-bold">{getMonthName(selectedMonth)}</h2>
        <Button label="Next" onClick={() => changeMonth('next')} />
    </div>

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

