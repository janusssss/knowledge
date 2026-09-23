const { GLib } = imports.gi;
const Main = imports.ui.main;
const St = imports.gi.St;

let label, timeoutId;

function getGpuUsage() {
    try {
        let [success, output] = GLib.spawn_command_line_sync('nvidia-smi --query-gpu=utilization.gpu --format=csv,noheader,nounits');
        if (success) {
            return output.toString().trim();
        }
    } catch (e) {
        logError(e, 'Failed to execute nvidia-smi');
    }
    return 'N/A';
}

function updateLabel() {
    let gpuUsage = getGpuUsage();
    label.set_text(`GPU: ${gpuUsage}%`);
}

function init() {
    label = new St.Label({ text: 'GPU: N/A', y_expand: true, y_align: St.Align.MIDDLE });
}

function enable() {
    // 使用官方API添加到状态区域
    Main.panel.addToStatusArea('gpu-monitor-indicator', label);

    // 设置定时器每秒更新一次
    timeoutId = GLib.timeout_add(GLib.PRIORITY_DEFAULT, 1000, () => {
        updateLabel();
        return true; // 返回true表示继续定时器
    });
}

function disable() {
    if (label) {
        label.destroy();
    }
    if (timeoutId) {
        GLib.source_remove(timeoutId);
    }
}