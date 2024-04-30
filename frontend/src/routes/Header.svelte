<script>
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { BlameStartup } from "$lib/wailsjs/go/main/App";
	import { Button, Modal } from 'flowbite-svelte';
	import { beforeUpdate } from 'svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';

	let startup = Boolean(false);
	let init = Boolean(false);
	let defaultModal = Boolean(false);
	let dis = Boolean(false);
	let test = Boolean(false);
	let modalText = "Loading assets...";
	if (init == false) {
		defaultModal = true;
		BlameStartup().then((result) => (startup = result));
		init = true;

	}


	$:if (startup == true) {
		modalText = "Done Loading :)";
		setTimeout(modulo, 10000);
		function modulo() {
			defaultModal = false
	}
		
	}
</script>

<header>
	<nav class="navbar4">
		<ul>
		  <li class="new" aria-current={$page.url.pathname === '/' ? 'page' : undefined}><a href="/">Home</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/stueckliste' ? 'page' : undefined}><a href="/stueckliste">Stueckliste</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/verbindungsliste' ? 'page' : undefined}><a href="/verbindungsliste">verbindungsliste</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/filewatcher' ? 'page' : undefined}><a href="/filewatcher">filewatcher</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/cleaner' ? 'page' : undefined}><a href="/cleaner">Cleaner</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/test' ? 'page' : undefined}><a href="/test">Test</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/debugger' ? 'page' : undefined}><a href="/debugger">Debugger</a></li>
		</ul>
	  </nav>
</header>

  <Modal id="loading" bind:open={defaultModal} size="xs" autoclose bind:dismissable={dis}>
	<div class="text-center">
	  <ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" />
	  <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">{modalText}</h3>

	</div>
  </Modal>
<style>
	header {
		display: flex;
		justify-content: space-between;
	}

	nav {
		display: flex;
		justify-content: center;
		--background: rgba(255, 255, 255, 0.7);
		width: 100%;
	}

	ul {
		position: relative;
		padding: 0;
		margin: 0;
		height: 3em;

		display: flex;
		justify-content: center;
		align-items: center;
		list-style: none;
		background: var(--background);
		background-size: contain;
	}

	li {
		position: relative;
		height: 100%;
	}

	li[aria-current='page']::before {
		--size: 6px;
		content: '';
		width: 0;
		height: 0;
		position: absolute;
		top: 0;
		left: calc(50% - var(--size));
		border: var(--size) solid transparent;
		border-top: var(--size) solid var(--color-theme-1);
	}

	a {
		display: flex;
		height: 100%;
		align-items: center;
		padding: 0 0.5rem;
		color: var(--color-text);
		font-weight: 700;
		font-size: 0.8rem;
		text-transform: uppercase;
		letter-spacing: 0.1em;
		text-decoration: none;
		transition: color 0.2s linear;
		cursor: pointer;
	}

	a:hover {
		color: var(--color-theme-1);
	}
</style>
