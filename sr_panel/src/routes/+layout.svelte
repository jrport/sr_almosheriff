<script>
	import './app.css'
	import SideBar from './SideBar.svelte'
	import SearchBar from './SearchBar.svelte';
</script>

<div class="bg-cyan-800 flex h-screen w-screen flex-row">
	<SideBar/>
	<div class="flex flex-col w-full">
		<SearchBar/>
		<slot></slot>
	</div>
</div>
