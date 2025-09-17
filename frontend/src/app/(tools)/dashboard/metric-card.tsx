"use client";

import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardDescription,
} from "@/components/ui/card";

export default function MetricCard({
  title,
  description,
  data,
}: {
  title: string;
  description: string;
  data: Record<string, string>;
}) {
  return (
    <Card className="bg-secondary-background text-foreground">
      <CardHeader>
        <CardTitle className="text-base">{title}</CardTitle>
        <CardDescription>{description}</CardDescription>
      </CardHeader>
      <CardContent className="text-sm text-gray-800 font-bold">
        <ul>
          {Object.entries(data).map(([key, value]) => (
            <li key={key}>
              {key}: {value}
            </li>
          ))}
        </ul>
      </CardContent>
    </Card>
  );
}
