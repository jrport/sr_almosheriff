<script>
	const titulos = ["Pedidos", 'Compras', 'Inventários', 'Performances']
	const urls = ["/pedidos", '/compras', '/inventarios', 'performances']
</script>

<aside class="shadow-2xl text-white text-xl py-4 bg-sky-950 w-2/12 flex flex-col items-center">
	<a class="py-3 w-9/12" href="/"><img src="./logo.svg"  alt="Logo da São Roque"></a>
	<ul class="flex flex-col py-3">
		{#each titulos as titulo,index}
			<li class="self-center transition ease-in-out delay-150 hover:-translate-y-1 hover:scale-125 duration-300">
				<a class="transition ease-in-out delay-150 hover:-translate-y-1 hover:scale-125 duration-300" href="{urls[index]}">{titulo}</a>
			</li>
		{/each}
	</ul>
	<small class="self-start absolute bottom-0 px-3 py-3">V-1.0</small>
</aside>
