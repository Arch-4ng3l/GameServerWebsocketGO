<script>
    import { Router, Route, Link} from "svelte-routing"
    import { onMount } from "svelte";
    import { auth, getCookie, setCookie } from "./store";
    import ObjectList from "./ObjectList.svelte";
    import ConnList from "./ConnList.svelte";
    import Signin from "./Signin.svelte";
    import Assets from "./Assets.svelte";
    import Logs from "./Logs.svelte";
    let socket;
    let url = "/api/auth";

    onMount(() => {

        socket = new WebSocket("ws://localhost:3000/api/");
        let token = getCookie("token");
        if (token === "") {
            return;
        }
        fetch(url, {
            method: "POST",
            body: JSON.stringify({
                token: token,
            }),
        })
        .then(response => {
            if(response.status === 200) {
                $auth = true;
            }
        })
    });

    function logOut() {
        $auth = false;
        setCookie("token", "");
        socket.close();
    }

</script>

<main>
    <Router>
    {#if $auth}
        <div class="navbar">
            <nav>
                <Link to="/objects">
                <div class="link">
                    <h3>Object</h3>
                </div>
                </Link>

                <Link to="/conn">
                <div class="link">
                    <h3>Connections</h3>
                </div>
                </Link>

                <Link to="/assets">
                <div class="link">
                    <h3>Assets</h3>
                </div>
                </Link>

                <Link to="/logs">
                <div class="link">
                    <h3>Logs</h3>
                </div>
                </Link>

            </nav>
            <br>
            <button on:click={logOut} class="link">Log Out</button>

        </div>

        <h1>Server Dashboard</h1>

        <Route path="/objects" >
            <ObjectList {socket}/>
        </Route>

        <Route path="/conn" component={ConnList}/>
        <Route path="/assets" component={Assets}/>
        <Route path="/logs" component={Logs}/>

    {:else}
        <nav>
            <Link to="/signin">
            <div class="link">
                <h2>Sign In</h2>
            </div>
            </Link>
        </nav>

        <Route path="/signin">

            <br>
            <br>
            <Signin/>
        </Route>

    {/if}
    </Router>
</main>

<style>

	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}

    .navbar{
        position: fixed;
        display: flexbox;
        top: 30%;
        width: 300px;
        height: auto;
        padding: 0;
        margin: 0;
        z-index: 9;
        overflow: hidden;
        box-shadow: 0 8px 30px 0 rgba(0,0,0,0.3);
        text-decoration: none;
    }
    .link {
        text-transform: uppercase;
        border-radius: 3px;
        border-width: 1px;
        border-style: solid;
        color: black;
        font-size: 2em;
        text-decoration: none;
        transition: ease-in-out 0.1s;
    }
    .link:hover{
        text-decoration: none;
        border-width: 3px;
		color: #ff3e00;
        font-weight: bold;
    }
</style>
