# 1. Create a new web service
Navigate to https://dashboard.render.com/select-repo?type=web and enter your repo you want to deploy.  
In this example we will be using https://github.com/fenny/fiber-hello-world.  
![](https://i.imgur.com/MyPUCm0.png)

# 2. Settings
**Name**: `Fiber`  
**Environment**: `Go`  
**Branch**: `master`  
**Build Command**: `go build main.go`  
**Start Command**: `./app`  
![](https://i.imgur.com/IpvfYGo.png)

# 3. Deploy service
When deploying your service, a terminal will show up.  
![](https://i.imgur.com/rcfAbSF.png)  

Navigate to your app url: https://fiber.onrender.com/  
And you should see `Hello, World!`
