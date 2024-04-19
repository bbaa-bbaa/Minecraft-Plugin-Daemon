package plugin

import (
	"fmt"
	"strings"

	"cgit.bbaa.fun/bbaa/minecraft-plugin-server/core/plugin/pluginabi"
	"github.com/samber/lo"
)

type MinecraftPosition struct {
	Position  [3]float64
	Dimension string
}

type TeleportCore struct {
	BasePlugin
}

func (tc *TeleportCore) Init(pm pluginabi.PluginManager) error {
	tc.BasePlugin.Init(pm, tc)
	return nil
}

func (tc *TeleportCore) TeleportPlayer(src string, dst string) error {
	// 先尝试直接 tp
	res := tc.RunCommand(fmt.Sprintf(`tp %s %s`, src, dst))
	if strings.Contains(res, "No entity") {
		return fmt.Errorf("目标玩家不存在")
	}
	if !strings.Contains(res, "Teleported") {
		// 跨世界 TP
		dstPi, err := tc.GetPlayerInfo(dst)
		if err != nil {
			return fmt.Errorf("无法获取目标玩家信息")
		}
		tc.RunCommand(fmt.Sprintf(`execute as %s rotated as %s in %s run tp %s`, src, dst, dstPi.LastLocation.Dimension, dst))
	}
	return nil
}

func (tc *TeleportCore) TeleportPosition(src string, dst MinecraftPosition) error {
	tc.RunCommand(fmt.Sprintf(`execute as %s rotated as %s in %s run tp %s`, src, src, dst.Dimension, strings.Join(lo.Map(dst.Position[:], func(item float64, index int) string {
		return fmt.Sprintf("%f", item)
	}), " ")))
	return nil
}

func (tc *TeleportCore) Teleport(src string, dst any) error {
	// 存储目前位置
	pi, err := tc.GetPlayerInfo(src)
	if err != nil {
		return err
	}
	pi.LastLocation = pi.Location
	pi.Commit()
	switch dst := dst.(type) {
	case string:
		return tc.TeleportPlayer(src, dst)
	case MinecraftPosition:
		return tc.TeleportPosition(src, dst)
	}
	return nil
}

func (tc *TeleportCore) Name() string {
	return "TeleportCore"
}

func (tc *TeleportCore) DisplayName() string {
	return "传送内核"
}