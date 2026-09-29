{#-
  normal_entry 单列版。

  内置模板走 lib.typ 的 regular-entry()，那里写死了
  grid(columns: (1fr, entries-date-and-location-width))，所以照抄内置那套就一定是
  两列。这里不复用它，直接把内容整块铺满，第二列整个丢掉。
  其他 entry 类型继续走内置模板，保持两列。

  用法：design.templates.normal_entry.main_column 里把 LOCATION、DATE 都写进去，
  date_and_location_column 这个字段会被忽略。
#}
#metadata("skip-content-area")

#context {
  let config = rendercv-config.get()
  let typography-line-spacing = config.at("typography-line-spacing")
  let entries-side-space = config.at("entries-side-space")

  set par(
    spacing: typography-line-spacing,
    leading: typography-line-spacing,
    justify: config.at("justify"),
  )
  set align(config.at("start-align"))
  set list(
    marker: (
      config.at("entries-highlights-bullet"),
      config.at("entries-highlights-nested-bullet"),
    ),
    indent: config.at("entries-highlights-space-left"),
    spacing: config.at("entries-highlights-space-between-items") + typography-line-spacing,
    body-indent: config.at("entries-highlights-space-between-bullet-and-text"),
  )

  block(
    [
{% for line in entry.main_column.splitlines() %}
      {{ line|indent(6) }}

{% endfor %}
    ],
    breakable: config.at("entries-allow-page-break"),
    below: config.at("sections-space-between-regular-entries") + typography-line-spacing,
    inset: (
      left: entries-side-space,
      right: entries-side-space,
    ),
    width: 100%,
  )
}
