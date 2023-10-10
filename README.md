# Virtual Try-On System

基于深度学习的虚拟试衣系统

**第十四届中国服务外包创新大赛获奖作品**

联系方式：huangwenxuan271828@163.com

项目链接：https://github.com/T4t4KAU/VirtualTryOnSystem.git

论文链接：https://ieeexplore.ieee.org/document/10241865

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/0.png?raw=true" alt="image1.png" style="width:100%; height:auto;">

## 推荐硬件

内存: 32G

CPU: 16核

尚未支持GPU

## 安装方法

首先须要拉取所需的docker镜像，包含了所有的预处理和推理模块

```bash
docker pull nephilimboy/openpose-cpu
docker pull venuns/parse-agnostic
docker pull venuns/vitons
docker pull venuns/humanparse
docker pull venuns/densepose
docker pull venuns/resize:viton
docker pull venuns/clothmask
docker pull venuns/super-resolution
```

在项目根目录下执行：

```
bash run.sh
```

启动单个实例，默认启动在8888端口，启动时间可能会较长

启动3组实例：

```
bash govton-start.sh
```

POST请求接口：

```
curl -X POST -F "image=@./image/test.jpg" -F "cloth=@./cloth/test.jpg" 127.0.0.1:8888/vitons
```

`./image/test.jpg`是一个示例，实际上就是人物图片的路径，`./cloth/test.jpg`是衣服图片的路径

如果请求成功的话 收到的HTTP响应中会包含试衣结果，等待的时间可能会有3到5分钟

启动后，可以访问测试页面:http://127.0.0.1:8888/index

这个页面仅仅用来测试

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/1.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

点击上述按钮可以进行人物和衣服的上传

## 系统架构

### 预处理模块
#### DensePose UV Parts Segmentation

DensePose 是一种计算机视觉技术，旨在将图像中的人体姿态估计与密集表面估计结合起 来，能够精确地将人体图像中的每个像素点与一个身体部位或身体表面相关联。

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/6.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

#### Human Keypoint Detection using Openpose

OpenPose 是一种计算机视觉技术，能够实时地检测图像或视频中人体的姿势和动作，以 C++和 Python 编写，基于深度学习算法。

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/7.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

#### Human Parsing via Part Grouping Network (PGN)

Human Parsing via Part Grouping Network (PGN)是一种图像语义分割技术，主要用于 分割人体部位和解析人体姿势。

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/8.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

#### Clothing-Agnostic Representation

为了删除原始图像中的服装细节，本项目引入了一个新的特征服装不可知表示(Cloth Agnostic Processing)，它使用姿势信息和分割图，彻底消除对原衣服的依赖，保留需要 复制的身体部位。

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/9.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

#### Cloth Mask Extraction via UNet

Cloth Mask Extraction via UNet 是一种基于深度学习的图像分割技术，主要用于从服装图像中提取出衣服的遮罩。

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/10.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

#### High Resolution via SRGAN

本项目额外引入的一种基于 SRGAN 的超分辨率算法，是对预处理结果的进一步优化，为本 团队基于项目需求的创新，能够显著提高最终的结果的效果。

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/11.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

### 推理模块

基于上述的预处理结果，应用高分辨率虚拟试衣算法生成最终的试衣图像，该算法旨在解决传统虚拟试穿技术中存在的对齐和遮挡问题

整个框架可以分为两个部分： 
1. Try-on Condition Generator：该生成器使用无关衣服的人体表示和衣服图案，同 时生成衣服的分割图和人体的变形。在这个过程中，生成器不会产生任何错位或像 素挤压的伪影。具体地说，该生成器使用无关衣服的人体表示和衣服图像作为输入 ，生成衣服的分割图和人体的变形。生成器的设计能够确保衣服图案与人体之间的 错位不会产生任何伪影，同时可以避免像素挤压的情况。这种设计可以提高虚拟试 穿的质量，使试穿效果更加真实和自然。 
2. Image Condition Generator: 试穿图像生成器使用这些输出来合成最终的试穿结 果图像，将衣服和人体进行融合，生成虚拟试穿图像。这个过程可以通过深度学习 模型来实现，具有很好的试穿效果和图像质量

架构图：


总体流程：

### 服务端

服务端的职责是启动、使用和调度已经容器化的计算模块。我们使用Go语言中的Go-Zero框架进行开发，利用事先定义的 protobuf 文件可以方便快捷地生成代码，也便于后期的扩展与更新。

服务端使用 RPC(基于gRPC框架) 与计算模块进行通信，具体应用中，对于占用较多计算资源的模块(如HumanParse)，在服务启动时就预先加载(热启动)并且常驻后台运行并监听请求，在运行时尽量不和其他计算任务同时进行。
对于占用较少计算资源模块(如ClothMask), 该模块不会预先启动或常驻运行，仅在需要时运行一次，避免白白浪费内存(将内存资源尽量留给HumanParse)，这样的模块将使用命令行和SDK直接启停。

服务端会用一个 pending 表记录正在运行的计算模块，如果有长作业在运行，那么不会调度其他的复杂计算任务与其同时运行(避免为机器带来过高的负载而导致崩溃)。
对于短期且资源需求较低的计算任务，可以调度其他计算任务运行。服务端也不会调度多个复杂计算任务连续执行，以免简单计算任务发生"饥饿"。
有些计算任务之间具有依赖性，前一个任务的结果会是后一个任务的输入，服务端会将一些计算任务的输出结果暂存到队列中，以待相关任务被调度时使用。

但是对于内存小于32G的机器不推荐启动多个实例，建议每个容器只启动一个实例。

目录结构：
```
.
├── README.md
├── containers  # 容器操作实现
│   ├── agnostic
│   ├── containers.go
│   ├── densepose
│   ├── mask
│   ├── openpose
│   ├── parsing
│   ├── resize
│   ├── resolution
│   └── vitons
├── core
│   ├── etc
│   ├── govton.go
│   └── internal
│       ├── config
│       │   └── config.go
│       ├── handler # 请求处理
│       ├── logic  # 业务逻辑
│       ├── svc
│       │   └── service_context.go
│       └── types
│           └── types.go
├── go.mod
├── go.sum
├── govton-restart.sh
├── govton-start.sh
├── govton-stop.sh
├── govton.api
├── nginx.conf
├── process  # 预处理调用
│   ├── consts.go
│   ├── densepose.go
│   ├── human_agnostic.go
│   ├── mask.go
│   ├── parse.go
│   ├── parse_agnostic.go
│   ├── proto  # proto生成代码
│   ├── viton.go
│   └── vitons.go
├── run.sh
├── static # 静态文件
├── test   # 测试代码
└── utils
    └── utils.go
```

### OverView

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/2.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

## 主要工作

封装和改进了所使用的深度学习算法，在不影响推理结果的基础上，对大量论文代码进行了工程化或复现。

将多个预处理模块和推理计算模块进行容器化，并为容器增加了 RPC 通信功能(基于gRPC)和日志输出，解决了深度学习项目的移植性难题，可以通过 Docker 将模块快速部署到Linux服务器上，且无需安装其他依赖。

开发了WEB服务端程序负责使用和调度预处理和推理计算模块，对外部暴露一个HTTP接口，便于用户进行图片的上传和结果的接收。使用 Nginx 对引用进行反向代理，及时记录访问记录并拒绝恶意的请求。

## 试衣效果

展示1：

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/3.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

展示2：

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/4.png?raw=true" alt="image1.png" style="width:50%; height:auto;">

展示3：

<img src="https://github.com/T4t4KAU/VirtualTryOnSystem/blob/main/static/5.png?raw=true" alt="image1.png" style="width:80%; height:auto;">

## 总结

我们的团队借助深度学习算法和网络技术，成功实现了虚拟试衣系统。该系统不仅可以更准确地模拟人体的形态和肤色，还可以为消费者提供更加直观的购物体验。消费者可以更轻松地了解自己在不同款式和颜色的衣服中的形象效果，从而做出更明智的购买决策。在未来，随着技术的不断创新和完善，我们相信虚拟试衣项目会助成一种智能化、现代化的购物方式。这将大大提高消费者的购物效率和便利性，进一步促进线上购物的发展。

尽管我们在项目实现过程中面临着时间和资源的种种限制，但是我们凭借着团队协作和不断试错的精神，克服了重重困难，最终完成了该系统的开发。虽然结果还有很多不足之处，但是我们相信，凭借这份热爱与坚持，继续深耕这个领域，必能取得更大的进展和突破。从各个层面上讲，这次经历对我们的团队成员来说是一次宝贵的磨砺和提升。我们深刻体会到了团队合作和持续创新的重要性，也更加热爱和探索技术的未来。我们将继续不断地学习和努力，为更好地服务用户和推动技术发展做出更大的贡献。

## Give a star! ⭐

If you think this project is interesting, or helpful to you, please give a star!
