package dal

import (
	"context"

	"github.com/pion/rtp"
	"github.com/zobinHuang/BrosCloud/provider/streamer/model"
)

/*
	@struct: WebRTCStreamDAL
	@description: DAL layer
*/
type WebRTCStreamDAL struct {
	StreamerMap map[string]*model.WebRTCStreamer
	PipeMap     map[string]*model.WebRTCPipe
}

/*
	@struct: WebRTCStreamDALConfig
	@description: used for config instance of struct WebRTCStreamDAL
*/
type WebRTCStreamDALConfig struct{}

/*
	@function: NewWebRTCStreamDAL
	@description:
		create, config and return an instance of struct WebRTCStreamDAL
*/
func NewWebRTCStreamDAL(c *WebRTCStreamDALConfig) model.WebRTCStreamDAL {
	// allocate space for global maps
	streamMap := make(map[string]*model.WebRTCStreamer)
	pipeMap := make(map[string]*model.WebRTCPipe)

	wrtcDal := &WebRTCStreamDAL{
		StreamerMap: streamMap,
		PipeMap:     pipeMap,
	}

	return wrtcDal
}

/*
	@function: NewWebRTCStreamer
	@description:
		create a new webrtc streamer for a new instance
*/
func (d *WebRTCStreamDAL) NewWebRTCStreamer(ctx context.Context, streamInstance *model.StreamInstanceDaemonModel) (*model.WebRTCStreamer, error) {
	// create RTP channel
	videoStreamChan := make(chan *rtp.Packet, 1)
	audioStreamChan := make(chan *rtp.Packet, 1)

	// create new webrtc streamer
	webrtcStreamer := &model.WebRTCStreamer{
		StreamInstance: streamInstance,
		VideoStream:    videoStreamChan,
		AudioStream:    audioStreamChan,
	}

	// insert to global map
	d.StreamerMap[streamInstance.Instanceid] = webrtcStreamer

	return webrtcStreamer, nil
}

/*
	@function: NewWebRTCPipe
	@description:
		create a new webrtc pipe for a new consumer
*/
func (d *WebRTCStreamDAL) NewWebRTCPipe(ctx context.Context, streamInstance *model.StreamInstanceDaemonModel, consumerID string) (*model.WebRTCPipe, error) {
	// allocate memory space for three channels of WebRTC pipe
	videoChan := make(chan *rtp.Packet, model.VIDEO_PIPE_CHANNEL_LENGTH)
	audioChan := make(chan *rtp.Packet, model.AUDIO_PIPE_CHANNEL_LENGTH)
	inputChan := make(chan []byte, model.INPUT_PIPE_CHANNEL_LENGTH)

	webRTCPipe := &model.WebRTCPipe{
		StreamInstance: streamInstance,
		ConsumerID:     consumerID,
		VideoChan:      videoChan,
		AudioChan:      audioChan,
		InputChan:      inputChan,
	}

	return webRTCPipe, nil
}