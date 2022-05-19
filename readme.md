# Docker builder

## 개요

M1 맥북에서 이미지를 빌드할 때마다 `docker build --platform=linux/amd64 -t <tag> .` 처럼 플랫폼을 배포용 클러스터에 맞춰야 했는데요, 빌드 시에 필요한 데이터를 정리해주는 대화형 CLI 툴을 제작하고자 합니다.

## 목표

**입력값**

- 도커파일 경로 (default : 현재 터미널 디렉토리) - input
- 타겟 아키텍처 정보 (default : amd64) - select
- 이미지 태그 (required) - input
