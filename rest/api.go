package rest

import (
  "github.com/gin-gonic/gin"
  "github.com/dev-op-spec/engine/core"
  "net/http"
  "github.com/dev-op-spec/engine/core/models"
  "time"
  "io"
  "fmt"
)

type Api interface {
  Start()
}

func New(
coreApi core.Api,
) Api {

  return &_api{
    compositionRoot:newCompositionRoot(coreApi),
  }

}

type _api struct {
  compositionRoot compositionRoot
}

func (api _api) Start(
) {

  router := gin.Default()

  router.GET("/stream", func(c *gin.Context) {

    ticker := time.NewTicker(3 * time.Second)
    listener := ticker.C

    defer func() {
      ticker.Stop()
    }()

    c.Stream(func(w io.Writer) bool {

      fmt.Println()
      c.SSEvent("message", <-listener)
      return true

    })
  })

  router.GET("/dev-ops", func(c *gin.Context) {

    devOps, err := api.compositionRoot.CoreApi().ListDevOps()
    if (nil != err) {
      panic(err)
    }

    c.JSON(http.StatusOK, devOps)

  })

  router.POST("/dev-ops", func(c *gin.Context) {

    var req models.AddDevOpReq

    c.BindJSON(&req)

    err := api.compositionRoot.CoreApi().AddDevOp(req)
    if (nil != err) {
      panic(err)
    }

    c.Status(http.StatusOK)

  })

  router.PUT("/dev-ops/:dev-op-name/description", func(c *gin.Context) {

    var description string

    c.BindJSON(&description)

    req := models.NewSetDescriptionOfDevOpReq(
      description,
      c.Param("dev-op-name"),
    )

    err := api.compositionRoot.CoreApi().SetDescriptionOfDevOp(*req)
    if (nil != err) {
      panic(err)
    }

    c.Status(http.StatusOK)

  })

  router.POST("/dev-ops/:dev-op-name/runs", func(c *gin.Context) {

    devOpRunView, err := api.compositionRoot.CoreApi().RunDevOp(c.Param("dev-op-name"))
    if (nil != err) {
      panic(err)
    }

    c.JSON(http.StatusOK, devOpRunView)

  })

  router.GET("/pipelines", func(c *gin.Context) {

    pipelines, err := api.compositionRoot.CoreApi().ListPipelines()
    if (nil != err) {
      panic(err)
    }

    c.JSON(http.StatusOK, pipelines)

  })

  router.POST("/pipelines", func(c *gin.Context) {

    var req models.AddPipelineReq

    c.BindJSON(&req)

    err := api.compositionRoot.CoreApi().AddPipeline(req)
    if (nil != err) {
      panic(err)
    }

    c.Status(http.StatusOK)

  })

  router.PUT("/pipelines/:pipeline-name/description", func(c *gin.Context) {

    var description string

    c.BindJSON(&description)

    req := models.NewSetDescriptionOfPipelineReq(
      description,
      c.Param("pipeline-name"),
    )

    err := api.compositionRoot.CoreApi().SetDescriptionOfPipeline(*req)
    if (nil != err) {
      panic(err)
    }

    c.Status(http.StatusOK)

  })

  router.POST("/pipelines/:pipeline-name/runs", func(c *gin.Context) {

    pipelineRunView, err := api.compositionRoot.CoreApi().RunPipeline(c.Param("pipeline-name"))
    if (nil != err) {
      panic(err)
    }

    c.JSON(http.StatusOK, pipelineRunView)

  })

  router.POST("/pipelines/:pipeline-name/stages", func(c *gin.Context) {

    var req models.AddStageToPipelineReq

    c.BindJSON(&req)

    req.PipelineName = c.Param("pipeline-name")

    err := api.compositionRoot.CoreApi().AddStageToPipeline(req)
    if (nil != err) {
      panic(err)
    }

    c.Status(http.StatusOK)

  })

  router.Run()

}
