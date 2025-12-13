"use client";

import { Cell, Label, Pie, PieChart } from "recharts";

import * as React from "react";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

export default function PopularPageChart({
  data,
}: {
  data: { pagePath: string; pageName: string; views: number }[];
}) {
  const chartConfig: ChartConfig = Object.fromEntries(
    data.map((item, index) => [
      item.pageName,
      {
        label: item.pageName,
        color: `var(--chart-${(index % 12) + 1})`, // cycling 12 color
      },
    ]),
  );

  const totalVisitors = React.useMemo(() => {
    return data.reduce((acc, curr) => acc + curr.views, 0);
  }, [data]);

  return (
    <Card className="flex flex-col bg-secondary-background text-foreground">
      <CardHeader className="items-center pb-0">
        <CardTitle>Frequently visited page</CardTitle>
        <CardDescription>
          Showing frequentyl visited page in the last 30 days
        </CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="mx-auto max-h-40">
          <PieChart>
            <ChartTooltip
              cursor={false}
              content={(props) => (
                <ChartTooltipContent {...props} hideLabel />
              )}
            />
            <Pie
              data={data}
              dataKey="views"
              nameKey="pageName"
              innerRadius={35}
              strokeWidth={2}
            >
              {data.map((entry, index) => (
                <Cell
                  key={`cell-${index}`}
                  fill={`var(--chart-${(index % 12) + 1})`}
                />
              ))}
              <Label
                content={({ viewBox }) => {
                  if (viewBox && "cx" in viewBox && "cy" in viewBox) {
                    return (
                      <text
                        x={viewBox.cx}
                        y={viewBox.cy}
                        textAnchor="middle"
                        dominantBaseline="middle"
                      >
                        <tspan
                          x={viewBox.cx}
                          y={viewBox.cy}
                          className="fill-foreground text-xl font-bold"
                        >
                          {totalVisitors.toLocaleString()}
                        </tspan>
                        <tspan
                          x={viewBox.cx}
                          y={(viewBox.cy || 0) + 19}
                          className="fill-foreground"
                        >
                          Visitors
                        </tspan>
                      </text>
                    );
                  }
                }}
              />
            </Pie>
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
