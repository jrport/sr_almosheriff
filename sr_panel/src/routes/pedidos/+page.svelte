<script>
	export let data
	const pedidos = data['pedidos']
</script>

<main class="h-screen w-full">
	<h1>Pedidos aqui</h1>
	{#each pedidos as pedido}
		 <!-- content here -->
		<p>{pedido.Item}</p>
		<p>{pedido.DataPedido}</p>
		<p>{pedido.Origem}</p>
		<p>{pedido.Quantidade}</p>
	{/each}
</main>
