<script lang="ts">
	import Sun from 'lucide-svelte/icons/sun';
	import Moon from 'lucide-svelte/icons/moon';

	import { toggleMode } from 'mode-watcher';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';

	let screen = $state('Select a screen');

	import * as Table from '$lib/components/ui/table/index.js';
	import { GetResourceParams, LocalSearch } from '$lib/wailsjs/go/app/Search';
	import type { schema } from '$lib/wailsjs/go/models';

	let params = $state<string[]>([]);
	let searchResponse = $state<schema.LocalSearchResponse>();

	function getParams(resource: string) {
		GetResourceParams(resource).then((paramsList) => {
			params = paramsList;
		});
	}

	// table data
	const invoices = [
		{
			invoice: 'INV001',
			paymentStatus: 'Paid',
			totalAmount: '$250.00',
			paymentMethod: 'Credit Card',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV002',
			paymentStatus: 'Pending',
			totalAmount: '$150.00',
			paymentMethod: 'PayPal',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV003',
			paymentStatus: 'Unpaid',
			totalAmount: '$350.00',
			paymentMethod: 'Bank Transfer',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV004',
			paymentStatus: 'Paid',
			totalAmount: '$450.00',
			paymentMethod: 'Credit Card',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV005',
			paymentStatus: 'Paid',
			totalAmount: '$550.00',
			paymentMethod: 'PayPal',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV006',
			paymentStatus: 'Pending',
			totalAmount: '$200.00',
			paymentMethod: 'Bank Transfer',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV007',
			paymentStatus: 'Unpaid',
			totalAmount: '$300.00',
			paymentMethod: 'Credit Card',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV003',
			paymentStatus: 'Unpaid',
			totalAmount: '$350.00',
			paymentMethod: 'Bank Transfer',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV004',
			paymentStatus: 'Paid',
			totalAmount: '$450.00',
			paymentMethod: 'Credit Card',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV005',
			paymentStatus: 'Paid',
			totalAmount: '$550.00',
			paymentMethod: 'PayPal',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV006',
			paymentStatus: 'Pending',
			totalAmount: '$200.00',
			paymentMethod: 'Bank Transfer',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV007',
			paymentStatus: 'Unpaid',
			totalAmount: '$300.00',
			paymentMethod: 'Credit Card',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV003',
			paymentStatus: 'Unpaid',
			totalAmount: '$350.00',
			paymentMethod: 'Bank Transfer',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV004',
			paymentStatus: 'Paid',
			totalAmount: '$450.00',
			paymentMethod: 'Credit Card',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV005',
			paymentStatus: 'Paid',
			totalAmount: '$550.00',
			paymentMethod: 'PayPal',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV006',
			paymentStatus: 'Pending',
			totalAmount: '$200.00',
			paymentMethod: 'Bank Transfer',
			name: 'open',
			game: 'priority'
		},
		{
			invoice: 'INV007',
			paymentStatus: 'Unpaid',
			totalAmount: '$300.00',
			paymentMethod: 'Credit Card',
			name: 'open',
			game: 'priority'
		}
	];

	let searchInput = $state('');
	let reaction = $state('');
	let timeout: number;

	var search = () => {
		let params = selectedParams.join(',');
		LocalSearch(searchInput, screen.toLowerCase(), '6310b48233cbac5151bc9594', params).then(
			(resp) => (searchResponse = resp)
		);
	};

	// Listen for input and search only when the user stops typing for 1 second or lenght of search term is greater than 3
	$effect(() => {
		if (searchInput.length > 3) {
			// Condition to check
			// Clear the previous timeout if it exists
			if (timeout) {
				clearTimeout(timeout);
				reaction = '';
			}

			// Set a new timeout
			timeout = setTimeout(() => {
				search();
			}, 1000); // 1 second delay
		} else {
			reaction = '';
		}
	});

	let selectedParams = $state<string[]>([]);
</script>

<div class="flex h-screen w-screen flex-col space-y-4 p-4">
	<div class="flex flex-row">
		<div class="flex-2 flex h-20 w-full items-center justify-start">
			<h1 class="scroll-m-20 text-4xl font-bold tracking-tight lg:text-5xl">
				Search Demoooo {reaction}
			</h1>
		</div>
		<Button onclick={toggleMode} variant="outline" size="icon">
			<Sun
				class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
			<Moon
				class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
			<span class="sr-only">Toggle theme</span>
		</Button>
	</div>
	<div class="flex-8 flex h-full w-full flex-row items-center justify-center space-x-4">
		<div class="flex-4 sb-5 flex h-full w-44 flex-col space-y-6 pt-8">
			<Button
				onclick={() => {
					screen = 'Order';
					getParams('order');
				}}
				variant="secondary"
				class="flex h-16 flex-initial">Order</Button>
			<Button
				onclick={() => {
					screen = 'ARN';
					getParams('arn');
				}}
				variant="secondary"
				class="flex h-16 flex-initial">ARN</Button>
		</div>
		<div class="flex-6 border-white-500 flex h-full w-full flex-col space-y-2 border p-2">
			<div class="flex h-12 flex-initial items-center justify-start pl-8">
				<h2
					class="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
					{screen}
				</h2>
				<h4 class="flex flex-row items-center justify-center pl-8">
					{#if selectedParams.length > 0}
						<p class="ml-4">Search on:</p>
					{/if}
					{#each selectedParams as param, i (i)}
						<Button variant="outline" class="m-1 h-7">{param}</Button>
					{/each}
				</h4>
			</div>
			<div class="flex h-12 flex-initial flex-row">
				<div class="flex w-2/4 flex-auto items-center justify-center">
					<p class="flex flex-initial px-8">Search</p>

					<Input
						bind:value={searchInput}
						type="email"
						placeholder="type anything..."
						class="flex-intial flex max-w-xs" />
				</div>
				<div class="flex w-full flex-auto flex-row">
					<p class="flex-start flex items-center justify-center px-6">Params:</p>

					<div class="flex w-96 flex-auto items-center justify-start overflow-x-auto">
						{#each params as param, i (i)}
							<Button
								onclick={() => {
									// Get index of param
									let paramIndex = selectedParams.indexOf(param);
									if (paramIndex === -1) {
										// doesn't exist
										selectedParams.push(param);
									} else {
										selectedParams.splice(paramIndex, 1);
									}
								}}
								class="mr-2 h-7">{param}</Button>
						{/each}
					</div>
				</div>
			</div>
			<div class="h-96 flex-auto overflow-y-auto pt-4">
				<Table.Root>
					<Table.Header>
						<Table.Row>
							<Table.Head class="w-[100px]">Client</Table.Head>
							<Table.Head>Order #</Table.Head>
							<Table.Head>Customer Name</Table.Head>
							<Table.Head>Order Type</Table.Head>
							<Table.Head class="text-right">Status</Table.Head>
						</Table.Row>
					</Table.Header>
					<Table.Body>
						{#if searchResponse}
							{#each searchResponse.payload as result, i (i)}
								<Table.Row>
									<Table.Cell class="font-medium">{result.client_name}</Table.Cell>
									<Table.Cell>{result.code}</Table.Cell>
									<Table.Cell>{result.shipping_address?.name}</Table.Cell>
									<Table.Cell>{result.order_type}</Table.Cell>
									<Table.Cell class="text-right">{result.status}</Table.Cell>
								</Table.Row>
							{/each}
						{/if}
					</Table.Body>
				</Table.Root>
			</div>
		</div>
	</div>
</div>
