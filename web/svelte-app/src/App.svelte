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
        if (socket === undefined) {
            socket = new WebSocket("ws://localhost:3000/api");
        } else if (socket.ReadyState !== 1) {
            socket = new WebSocket("ws://localhost:3000/api");
        }

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
        console.log("Log Out");
        $auth = false;
        setCookie("token", "");
        try {
            socket.close();
        }
        catch {

        }
        window.location.href = "/";
    }

</script>

<main>
{#if $auth}
    <Router>
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

                <div class="link" on:click={logOut} on:keypress={logOut}> 
                    <h3>Log Out</h3>
                </div>
            </nav>
            <br>

        </div>

        <h1>Server Dashboard</h1>

        <Route path="/objects" >
            <ObjectList {socket}/>
        </Route>

        <Route path="/conn" component={ConnList}/>
        <Route path="/assets" component={Assets}/>
        <Route path="/logs" component={Logs}/>

    </Router>
{:else}
    <Signin/>

{/if}
    
</main>

<style>

    .navbar {
        border-radius: 10px;
        overflow: hidden;
        background-color: #333;
    }
    .link {
        float: left;
        color: #f2f2f2;
        text-align: center;
        padding: 14px 16px;
        text-decoration: none;
        font-size: 17px;
    }

    .link:hover{
        background-color: #ddd;
        color: black;
    }

    .link:active {
        background-color: #ff3e00;
        color: white;
    }

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

</style>
