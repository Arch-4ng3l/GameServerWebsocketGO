<script>
    import { auth, setCookie } from "./store";

    function signIn() {
       const username = document.getElementById("un").value;
       const password = document.getElementById("pw").value; 
       const obj = {
            username: username, 
            password: password,
       };

       const jsonObj = JSON.stringify(obj);
       let url = "/api/login"
       fetch(url, {
            method: "POST",
            body: jsonObj,
       })
       .then(response => {
            if(response.status === 200) {
                return response.json();
            }
       })
       .then(data => {
            const jwttoken = data.token;
            setCookie("token", jwttoken);
            $auth = true;
       });
        
    }
</script>


<div>
    <h1>Login</h1>
    <br>
    <br>
    <div>
        <label for="username"><b>Username</b></label> 
        <input class="in" id="un" type="text" name="username"/>
        <br>
        <br>
        <label for="passwd"><b>Password</b></label>
        <input class="in" id="pw" type="password" name="passwd"/>
    </div>
    <br>
    <br>

    <button on:click={signIn} class="btn"><b>Sign In</b></button>
</div>


<style>
    div { 
        align-items: center;
    }
    .in {
        border-radius: 10px;
    }
    .in:hover {
        border-color: #ff3e00;
    }
    .btn {
        width: 10%;
        background: #454a4d;
        border-radius: 10px;     
    }
    .btn:hover {
        border-color: #ff3e00;
        color: #ff3e00;
        font-weight: bold;
    }
</style>
