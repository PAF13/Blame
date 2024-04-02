<script>
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	function GotoHome() {
		goto('/stueckliste');
	}
	function GotoStuecklsite() {
		goto('/stueckliste');
	}
	function GotoVerbindungliste(){
		goto('/stueckliste');
	}
	function GotoFilewatcher(){
		goto('/stueckliste');
	}
</script>

<header>
	<nav class="navbar4">
		<ul>
		  <li class="new" aria-current={$page.url.pathname === '/' ? 'page' : undefined}><a href="/">Home</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/stueckliste' ? 'page' : undefined}><a href="/stueckliste">Stueckliste</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/verbindungsliste' ? 'page' : undefined}><a href="/verbindungsliste">verbindungsliste</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/filewatcher' ? 'page' : undefined}><a href="/filewatcher">filewatcher</a></li>
		  <li class="new" aria-current={$page.url.pathname === '/test' ? 'page' : undefined}><a href="/test">Test</a></li>
		</ul>
	  </nav>
</header>

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
